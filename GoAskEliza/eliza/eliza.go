package eliza

import (
    "./answers/generator"
    "./answers/picker"
)

type Eliza struct {
    // Eliza struct can be created using 2
    // interfaces, this allows you to provide multiple
    // different implementations of how you generate answers
    // and how you pick them.
    generator generator.AnswerGenerator
    picker picker.AnswerPicker
}

func NewEliza(generator generator.AnswerGenerator, picker picker.AnswerPicker) *Eliza {
    eliza := Eliza{generator:generator, picker:picker}
    return &eliza
} 

func (e Eliza) GoAsk(question string) string {
    answers := e.generator.GenerateAnswers(question)
    return e.picker.PickAnswer(answers)
} 
