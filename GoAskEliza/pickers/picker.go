package pickers


type RandomPicker struct {

}

func (picker RandomPicker) PickAnswer(answers []string) string{
    return answers[0]
}