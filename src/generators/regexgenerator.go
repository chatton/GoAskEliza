package generators

import (
	// package used for regular expressions.
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

type RegexGenerator struct {
	// unexported response map of regular expressions to list of answers.
	responseMap   map[*regexp.Regexp][]string
	reflectionMap map[string]string
	pastQuestions map[string]bool
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
	generator.pastQuestions = make(map[string]bool)
	return generator
}

func (gen RegexGenerator) isRepeatQuestion(question string) bool {
	return gen.pastQuestions[question]
}

func (gen RegexGenerator) rememberQuestion(question string) {
	gen.pastQuestions[question] = true
}

func repeatAnswers() []string {
	return []string{
		"Hmmm, you've asked this before.",
		"I see you want to talk about this some more.",
		"It's interesting that you want to talk about this again.",
		"I find it interesting that you're talking about this again."}
}

// function to dig up a past question so that it can be used in
// a question when no other response is better.
func (gen RegexGenerator) getRandomPastQuestion() string {
	rand.Seed(time.Now().UTC().UnixNano())
	i := rand.Intn(len(gen.pastQuestions))
	count := 0
	for question := range gen.pastQuestions {
		if i == count {
			return question
		}
		count++
	}

	panic("should not be possible to reach here.")
}

func (gen RegexGenerator) GenerateAnswers(question string) []string {
	question = strings.ToLower(question) // ignore case
	if gen.isRepeatQuestion(question) {
		// if they ask the same question, they will get a response showing
		// that the previous question was "remembered"
		return repeatAnswers()
	}

	// question will now prompt repeat answers if it comes up again
	gen.rememberQuestion(question)

	for re, possibleResponses := range gen.responseMap {
		if re.MatchString(question) {
			questionTopic := gen.getQuestionTopic(re, question)
			returnResponses := []string{}
			for _, response := range possibleResponses {
				// it means the response will use the question topic in the answer and needs to be formatted.
				if strings.Contains(response, "%s") {
					// insert the question topic into the response.
					returnResponses = append(returnResponses, fmt.Sprintf(response, questionTopic))
				} else {
					// the response doesn't need to be formatted. It is complete as is.
					returnResponses = append(returnResponses, response)
				}
			}
			return returnResponses
		}
	}
	// no match was found, repsond with generic answers.
	return gen.defaultAnswers()
}

func (gen RegexGenerator) defaultAnswers() []string {
	// provide some answers in the case of no regex match on the question.

	// question that makes use of a random past question the user asked.
	// intended to make the responses seem more like a real life conversation.
	reflectOnPreviousQuestion := fmt.Sprintf("You asked \"%s\", let's talk some more about that.",
		gen.getRandomPastQuestion())

	// some generic catch all answers
	genericAnswers := []string{
		"I don't know how to respond to that",
		"Hmmm interesting...",
		"Tell me more.",
		"Please, continue.",
		"Could you elaborate on that?"}

	if len(gen.pastQuestions) > 0 { // there is at least one past question to dig up.
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
