package generator

// an answer generator should be able to give back a slice of answers when given a question.
type AnswerGenerator interface {
    GenerateAnswers(question string) []string
}

type RegexGenerator struct {}

func (r RegexGenerator) GenerateAnswers(question string) []string {
    return []string{ "One", "Two", "Three"}
}