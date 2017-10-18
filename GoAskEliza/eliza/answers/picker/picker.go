package picker

// answer picker is in charge of picking a response out of a list of answers.
type AnswerPicker interface {
    PickAnswer(answers []string) string
}

type RandomPicker struct {

}

func (picker RandomPicker) PickAnswer(answers []string) string{
    return answers[0]
}