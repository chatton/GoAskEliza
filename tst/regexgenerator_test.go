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

	// on a repeat question the repeat answers are given.
	testData = append(testData, TestData{"I like waffles.", []string{
		"Hmmm, you've asked this before.",
		"I see you want to talk about this some more.",
		"It's interesting that you want to talk about this again.",
		"I find it interesting that you're talking about this again.",
		"You seem to be repeating yourself.",
		"Are you expecting a different answer to the same question?",
	}})

	testData = append(testData, TestData{"I feel sad.", []string{
		"Can you tell me why you feel sad?",
		"Tell me more about these feelings.",
		"How long have you felt sad?",
	}})

	testData = append(testData, TestData{"I like talking to you.", []string{
		"Why do you like talking to me?",
		"Are you sure you like talking to me?",
	}})

	testData = append(testData, TestData{"I had a rough childhood.", []string{
		"What did you want to be when you were growing up?",
		"What was life like when you were growing up?",
	}})

	for _, data := range testData {
		answers := gen.GenerateAnswers(data.question)
		if !sliceContainsAll(answers, data.expectedAnswers) {
			t.Errorf("Test failed. Actual %s, Expected %s", answers, data.expectedAnswers)
		}
	}

}
