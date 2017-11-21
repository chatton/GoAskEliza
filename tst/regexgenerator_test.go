package test

import (
	"testing"

	"../src/generators"
)

type TestData struct {
	question        string
	expectedAnswers []string
}

func sliceContains(slice []string, val string) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}

func sliceContainsAll(actual, expected []string) bool {
	for _, expectedVal := range expected {
		if !sliceContains(actual, expectedVal) {
			return false
		}
	}
	return true
}

func TestRegexGeneratorResponses(t *testing.T) {
	gen := generators.NewRegexGenerator("./data/pattern-responses.dat")
	testData := []TestData{}

	// test first response that's not a hello gives back the rude responses.
	testData = append(testData, TestData{"Something rude.", []string{
		"Normally my clients start by saying 'hello'.",
		"No hello?",
		"I don't get a hello?",
		"You're not going to say hi?",
	}})

	testData = append(testData, TestData{"I like waffles.", []string{
		"Why do you like waffles?",
		"Are you sure you like waffles?",
	}})

	testData = append(testData, TestData{"My name is Bob.", []string{
		"Hello Bob, it's nice to meet you.",
		"It's nice to meet you, Bob.",
	}})

	for _, data := range testData {
		answers := gen.GenerateAnswers(data.question)
		if !sliceContainsAll(answers, data.expectedAnswers) {
			t.Errorf("Test failed. Actual %s, Expected %s", answers, data.expectedAnswers)
		}
	}

}
