package generators


import (
    // package used for regular expressions.
    "regexp"
    "os"
    "bufio"
    "strings"
)

type RegexGenerator struct {
    // unexported response map of regular expressions to list of answers.
    responseMap map[*regexp.Regexp] []string
}

func NewRegexGenerator(responsePatternPath string) RegexGenerator {
    generator := RegexGenerator{}
    // create the map of responses to possible answers.
    responseMap := makeResponseMap(responsePatternPath)
    generator.responseMap = responseMap
    return generator
}

func (r RegexGenerator) GenerateAnswers(question string) []string {
    return []string{ "One", "Two", "Three"}
}

// map used to map certain words from the question into an appropriate
// response in the answer.
var reflecionMap map[string] string = map[string]string{
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

var punctuation []string = []string{"!", ",", ";", ".", "?"}


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


