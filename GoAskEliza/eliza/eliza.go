package eliza


type Eliza struct {
    // Eliza struct can be created using 2
    // interfaces, this allows you to provide multiple
    // different implementations of how you generate answers
    // and how you pick them.
    generator AnswerGenerator
    picker AnswerPicker
}

func NewEliza(generator AnswerGenerator, picker AnswerPicker) *Eliza {
    eliza := Eliza{generator:generator, picker:picker}
    return &eliza
} 

func (e Eliza) GoAsk(question string) string {
    answers := e.generator.GenerateAnswers(question)
    return e.picker.PickAnswer(answers)
} 

// https://github.com/golang/go/wiki/CodeReviewComments#interfaces
// The documentation states that interfaces belong in the package that is
// going to use the interface type.

// an answer generator should be able to give back a slice of answers when given a question.
type AnswerGenerator interface {
    GenerateAnswers(question string) []string
}

// answer picker is in charge of picking a response out of a list of answers.
type AnswerPicker interface {
    PickAnswer(answers []string) string
}