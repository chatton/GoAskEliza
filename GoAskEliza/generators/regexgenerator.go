package generators


import (
    // package used for regular expressions.
    "regexp"
    "os"
    "bufio"
    "strings"
    "fmt"
)

type RegexGenerator struct {
    // unexported response map of regular expressions to list of answers.
    responseMap map[*regexp.Regexp] []string
    reflectionMap map[string] string
}

func NewRegexGenerator(responsePatternPath string) RegexGenerator {
    generator := RegexGenerator{}
    // create the map of responses to possible answers.
    responseMap := makeResponseMap(responsePatternPath)
    
    // map used to map certain words from the question into an appropriate
    // response in the answer.
   var reflectionMap map[string] string = map[string]string {
        "am": "are",
        "was": "were",
        "i": "you",
        "i'd": "you would",
        "i've": "you have",
        "i'll": "you will",
        "my": "your",
        "are": "am",
        "you've": "I have",
        "you'll": "I will",
        "your": "my",
        "yours": "mine",
        "you": "me",
        "me": "you",
    }

    generator.responseMap = responseMap
    generator.reflectionMap = reflectionMap
    return generator
}

func (gen RegexGenerator) GenerateAnswers(question string) []string {
    question = strings.ToLower(question) // ignore case
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
    return []string{"I don't know how to response to that."}
}

func (gen RegexGenerator) getQuestionTopic(re *regexp.Regexp, question string) string {
    match := re.FindStringSubmatch(question)
    questionTopic := match[1] // 0 is the full string, 1 is first match.
    questionTopic = gen.substituteWords(questionTopic)
    questionTopic = removePunctuation(questionTopic)
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
func removePunctuation(answer string) string {
    allChars := strings.Split(answer, "")
    for index, char := range allChars {
        for _, punc := range punctuation {
            if char == punc {
                allChars[index] = ""   
            }
        }
    }
    return strings.Join(allChars, "")
}


// include %s to strip incase the user enters "%s" directly into the question.
var punctuation []string = []string{"!", ",", ";", ".", "?", "%s"}


func makeResponseMap(path string) map[*regexp.Regexp] [] string {
    // map that will hold regex expressions and a list of possible responses
    // that will be read in from a file.
    resultMap := make(map[*regexp.Regexp] []string)
    file, err := os.Open(path)

    if err != nil { // something went wrong opening the file
        panic(err) // can't continue if the file isn't found.
    }

    defer file.Close() // close the file after this function finihes executing.

    // read the file line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() { // keep reading each line until we hit the end of the file.
        allPatterns := strings.Split(scanner.Text(), ";") // patterns on first line
        scanner.Scan() // responses on the next line
        allResponses := strings.Split(scanner.Text(), ";")
        for _, pattern := range allPatterns {
            pattern = "(?i)" + pattern // make every pattern case insensitive
            re := regexp.MustCompile(pattern) // throws an error if the pattern doesn't compile.
            resultMap[re] = allResponses
        }
    }
    return resultMap
}


