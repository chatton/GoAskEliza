package generators


type RegexGenerator struct {

}

func (r RegexGenerator) GenerateAnswers(question string) []string {
    return []string{ "One", "Two", "Three"}
}
