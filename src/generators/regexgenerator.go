package generators

import (
	// package used for regular expressions.
	"../util"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type RegexGenerator struct {
	// unexported response map of regular expressions to list of answers.
	responseMap   map[*regexp.Regexp][]string
	reflectionMap map[string]string
	pastQuestions *util.StringSet
}

// include %s to strip incase the user enters "%s" directly into the question.
var unwantedCharacters []string = []string{"!", ",", ";", ".", "?", "%s"}

func NewRegexGenerator(responsePatternPath string) RegexGenerator {
	generator := RegexGenerator{}
	// create the map of responses to possible answers.
	responseMap := makeResponseMap(responsePatternPath)

	// map used to map certain words from the question into an appropriate
	// response in the answer.
	var reflectionMap map[string]string = map[string]string{
		"am":     "are",
		"was":    "were",
		"i":      "you",
		"i'd":    "you would",
		"i've":   "you have",
		"i'll":   "you will",
		"my":     "your",
		"are":    "am",
		"you've": "I have",
		"you'll": "I will",
		"your":   "my",
		"yours":  "mine",
		"you":    "me",
		"me":     "you",
	}

	generator.responseMap = responseMap
	generator.reflectionMap = reflectionMap
	generator.pastQuestions = util.NewStringSet()
	return generator
}

func (gen RegexGenerator) isRepeatQuestion(question string) bool {
	return gen.pastQuestions.Contains(question)
}

func (gen RegexGenerator) rememberQuestion(question string) {
	gen.pastQuestions.Add(question)
}

func repeatAnswers() []string {
	// provide multiple generic responses for if the user asks a duplicate question.
	return []string{
		"Hmmm, you've asked this before.",
		"I see you want to talk about this some more.",
		"It's interesting that you want to talk about this again.",
		"I find it interesting that you're talking about this again.",
		"You seem to be repeating yourself.",
		"Are you expecting a different answer to the same question?"}
}

// function to dig up a past question so that it can be used in
// a question when no other response is better.
func (gen RegexGenerator) getRandomPastQuestion() string {
	return gen.pastQuestions.RandomValue()
}

func (gen RegexGenerator) GenerateAnswers(question string) []string {
	question = strings.ToLower(question) // ignore case
	if gen.isRepeatQuestion(question) {
		// if they ask the same question, they will get a response showing
		// that the previous question was "remembered"
		return repeatAnswers()
	}

	// don't want to include this question in "past" questions for the default answers.
	// so evaluate these now.
	defaultAnswers := gen.defaultAnswers(); 

	// question will now prompt repeat answers if it comes up again
	gen.rememberQuestion(question)

	for re, possibleResponses := range gen.responseMap {
		if re.MatchString(question) {
			questionTopic := gen.getQuestionTopic(re, question)
			returnResponses := []string{}
			for _, response := range possibleResponses {
				returnResponses = appendResponse(returnResponses, response, questionTopic)
			}
			return returnResponses
		}
	}
	// no match was found, repsond with generic answers.
	return defaultAnswers
}

func appendResponse(responses []string, response, questionTopic string) []string {
	// it means the response will use the question topic in the answer and needs to be formatted.
	if strings.Contains(response, "%s") {
		// insert the question topic into the response.
		return append(responses, fmt.Sprintf(response, questionTopic))
	} else {
		// the response doesn't need to be formatted. It is complete as is.
		return append(responses, response)
	}
}

func (gen RegexGenerator) defaultAnswers() []string {
	// provide some answers in the case of no regex match on the question.

	// some generic catch all answers
	genericAnswers := []string{
		"I don't know how to respond to that",
		"Hmmm interesting...",
		"Tell me more.",
		"Please, tell me more.",
		"Could you elaborate on that?"}

	if gen.pastQuestions.Size() > 0 { // there is at least one past question to dig up.
		// question that makes use of a random past question the user asked.
		// intended to make the responses seem more like a real life conversation.
		reflectOnPreviousQuestion := fmt.Sprintf("Earlier you said \"%s\", let's talk some more about that.",
			gen.getRandomPastQuestion())
		// give the chance that this will be brought up, not every time.
		genericAnswers = append(genericAnswers, reflectOnPreviousQuestion)
	}
	return genericAnswers
}

func (gen RegexGenerator) getQuestionTopic(re *regexp.Regexp, question string) string {
	match := re.FindStringSubmatch(question)
	questionTopic := match[1] // 0 is the full string, 1 is first match.
	questionTopic = gen.substituteWords(questionTopic)
	questionTopic = removeUnwantedCharacters(questionTopic)
	return questionTopic
}

func (gen RegexGenerator) substituteWords(answer string) string {
	allWords := strings.Split(answer, " ") // get slices of the words {"words", "in", "sentence"}
	for index, word := range allWords {
		// put to lower case so the capitilzation doesn't matter
		if val, ok := gen.reflectionMap[strings.ToLower(word)]; ok {
			allWords[index] = val // substitite the value
		}
	}
	return strings.Join(allWords, " ") // join back into string "words in sentence"
}

/*
removes punction from the string, this is to make it
so that the question returned doesn't contain unusual punctuation
taken from the input question.
*/
func removeUnwantedCharacters(answer string) string {
	for _, unwanted := range unwantedCharacters {
		// replace every unwanted string with an empty string
		// this implementation is not very efficent, O(n^m) to strip out the string.
		// m will always be very small, so this won't result in a huge performance hit.
		answer = strings.Replace(answer, unwanted, "", -1)
	}
	return answer

}

func makeResponseMap(path string) map[*regexp.Regexp][]string {
	// map that will hold regex expressions and a list of possible responses
	// that will be read in from a file.
	resultMap := make(map[*regexp.Regexp][]string)
	file, err := os.Open(path)

	if err != nil { // something went wrong opening the file
		panic(err) // can't continue if the file isn't found.
	}

	defer file.Close() // close the file after this function finishes executing.

	// read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // keep reading each line until we hit the end of the file.
		allPatterns := strings.Split(scanner.Text(), ";") // patterns on first line
		scanner.Scan()                                    // responses on the next line
		allResponses := strings.Split(scanner.Text(), ";")
		for _, pattern := range allPatterns {
			pattern = "(?i)" + pattern        // make every pattern case insensitive
			re := regexp.MustCompile(pattern) // throws an error if the pattern doesn't compile.
			resultMap[re] = allResponses
		}
	}
	return resultMap
}
