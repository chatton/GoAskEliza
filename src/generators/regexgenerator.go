package generators

import (
	"math/rand"

	"../util" // for the StringSet struct
	// for file IO
	"bufio"
	"fmt"
	// os for file management
	"os"
	// package used for regular expressions.
	"regexp"
	// for string manipulation
	"strings"
)

// A Response contains an regexp and then a slice of strings
// which are the possible responses once a question matches the pattern
// of the regexp.
type Response struct {
	// the regex pattern that matches the question
	re *regexp.Regexp
	// answers to that question
	responses []string
}

// RegexGenerator IS-AN answer generator. It makes use of regular
// expressions to match the pattern in a given question, and returns
// a list of possible responses.
type RegexGenerator struct {
	// no exported fields.
	responses     []Response
	reflectionMap map[string]string
	pastQuestions *util.StringSet
	firstQuestion bool
}

var unwantedCharacters []string
var genericAnswers []string
var rudeAnswers []string
var repeatAnswers []string
var greetingPatterns []string

func NewRegexGenerator(responsePatternPath string) *RegexGenerator {
	generator := &RegexGenerator{}
	generator.firstQuestion = true

	// store all these responses in memory. If you want to edit the files,
	// the program will need to be re-built and run again.
	// File IO is just done once, not every time the values are queried.
	unwantedCharacters = readLines("./data/unwanted.dat")
	genericAnswers = readLines("./data/generic-responses.dat")
	rudeAnswers = readLines("./data/rude-answers.dat")
	repeatAnswers = readLines("./data/repeat-answers.dat")
	greetingPatterns = makePatternsCaseInsensitive(readLines("./data/greeting-patterns.dat"))
	generator.responses = makeResponses(responsePatternPath)

	// map used to map certain words from the question into an appropriate
	// response in the answer.
	generator.reflectionMap = makeReflectionMap()
	generator.pastQuestions = util.NewStringSet()
	return generator
}

func makePatternsCaseInsensitive(patterns []string) []string {
	caseInsensitive := make([]string, len(patterns))
	for index, pattern := range patterns {
		caseInsensitive[index] = "(?i)" + pattern
	}
	return caseInsensitive
}

func makeReflectionMap() map[string]string {
	lines := readLines("./data/reflectionmap.dat")
	reflectionMap := make(map[string]string)
	for _, line := range lines {
		keyVal := strings.Split(line, ";") // {"i;you", "i'd;you would" ... }
		reflectionMap[keyVal[0]] = keyVal[1]
	}
	return reflectionMap
}

func getRandomElementFromSet(set *util.StringSet) string {
	values := set.Values()
	return values[rand.Intn(len(values))]
}

// function that will take in a file and give back a slice of strings, each
// holding one non-comment non-blank line.
func readLines(path string) []string {
	lines := []string{}
	file, err := os.Open(path)
	if err != nil { // something went wrong opening the file
		// fail fast - if a single file isn't found it needs to be fixed.
		panic(err) // can't continue if the file isn't found.
	}

	defer file.Close() // close the file after this function finishes executing.

	// read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if lineIsComment(line) { // skip comments.
			continue // by continuing, the scanner.Scan() statement in the loop will execute and skip this line of the file.
		}
		lines = append(lines, line)
	}

	return lines
}

func lineIsComment(line string) bool {
	return strings.HasPrefix(line, "#") || len(strings.TrimSpace(line)) == 0
}

func makeResponses(path string) []Response {
	allLines := readLines(path)      //read in slice of all lines in the file.
	responses := make([]Response, 0) // make a slice of Responses to hold the responses, don't know how many there will be so start at size = 0
	for i := 0; i < len(allLines); i += 2 {
		allPatterns := strings.Split(allLines[i], ";")    // patterns on first line
		allResponses := strings.Split(allLines[i+1], ";") // responses on the next line.
		for _, pattern := range allPatterns {
			pattern = "(?i)" + pattern        // make pattern case insensitive.
			re := regexp.MustCompile(pattern) // throws an error if the pattern doesn't compile.
			responses = append(responses, Response{re: re, responses: allResponses})
		}
	}
	return responses
}

func (gen *RegexGenerator) isRepeatQuestion(question string) bool {
	return gen.pastQuestions.Contains(question)
}

func (gen *RegexGenerator) rememberQuestion(question string) {
	gen.pastQuestions.Add(question)
}

// function to dig up a past question so that it can be used in
// a question when no other response is better.
func (gen *RegexGenerator) getRandomPastQuestion() string {
	return getRandomElementFromSet(gen.pastQuestions)
}

func questionIsGreeting(question string) bool {

	for _, pattern := range greetingPatterns {
		re := regexp.MustCompile(pattern)
		if re.MatchString(question) {
			return true
		}
	}
	return false
}

// GenerateAnswers belongs to eliza.AnswerGenerator interface.
func (gen *RegexGenerator) GenerateAnswers(question string) []string {

	// eliza will give back an answer recognizing that you didn't greet her.
	if gen.firstQuestion {
		gen.firstQuestion = false
		if !questionIsGreeting(question) {
			return rudeAnswers
		}
	}

	if gen.isRepeatQuestion(question) {
		// if they ask the same question, they will get a response showing
		// that the previous question was "remembered"
		return repeatAnswers
	}

	// don't want to include this question in "past" questions for the default answers.
	// so evaluate these now.
	defaultAnswers := gen.defaultAnswers()

	if !questionIsGreeting(question) { // don't want to "remember" a greeting.
		// end with with Eliza saying "Earlier you said "hello" let's talk more about that"
		// question will now prompt repeat answers if it comes up again
		gen.rememberQuestion(question)
	}

	for _, response := range gen.responses { // looking through all possible responses.
		re := response.re
		if re.MatchString(question) { // if the question matches a response regex pattern.

			// don't consider the element at index 0 which is the full match.
			matches := re.FindStringSubmatch(question)[1:]
			// only want to switch the pronouns and remove punctuation from the user input text.
			for index, val := range matches {
				matches[index] = gen.substituteWords(val)
				matches[index] = removeUnwantedCharacters(matches[index])
			}

			returnResponses := make([]string, len(response.responses)) // make our new slice to hold the returned responses.
			for index, resp := range response.responses {              // go through the possible return values for that response
				returnResponses[index] = constructAnswer(resp, matches)
			}
			return returnResponses // give back a slice of fully formed answers to the question.
		}
	}
	// no match was found, repsond with generic answers.
	return defaultAnswers
}

func constructAnswer(response string, matches []string) string {
	if len(matches) == 0 {
		return response // no substitution is requred, the answer is already complete.
	}
	// I consulted this SO question on how to call Sprintf using varargs
	//https://stackoverflow.com/questions/30588581/how-to-pass-variable-parameters-to-sprintf-in-golang
	istrings := make([]interface{}, len(matches))
	for index, val := range matches {
		istrings[index] = val
	}
	return fmt.Sprintf(response, istrings...)
}

func (gen *RegexGenerator) defaultAnswers() []string {
	// provide some answers in the case of no regex match on the question.

	returnAnswers := []string(genericAnswers) // don't want to modify the defaults. So we copy and add.
	if !gen.pastQuestions.IsEmpty() {         // there is at least one past question to dig up.
		// question that makes use of a random past question the user asked.
		// intended to make the responses seem more like a real life conversation.
		reflectOnPreviousQuestion := fmt.Sprintf("Earlier you said \"%s\", let's talk some more about that.",
			gen.getRandomPastQuestion())
		// give the chance that this will be brought up, not every time.
		returnAnswers = append(returnAnswers, reflectOnPreviousQuestion)
	}
	return returnAnswers
}

func (gen *RegexGenerator) substituteWords(answer string) string {
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
taken from the input question. Eg. "Why do you like waffles.?" should be
"Why do you like waffles?"
*/
func removeUnwantedCharacters(answer string) string {
	for _, unwanted := range unwantedCharacters {
		// replace every unwanted string with an empty string
		// this implementation is not very efficent, O(n^m) to strip out the string.
		// m will always be very small, so this won't result in a huge performance hit.
		answer = strings.Replace(answer, unwanted, "", -1) // -1 to specify all occurrences.
	}
	return answer
}
