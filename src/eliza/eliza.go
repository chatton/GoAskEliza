package eliza

type Eliza struct {
	// Eliza struct can be created using 2
	// interfaces, this allows you to provide multiple
	// different implementations of how you generate answers
	// and how you pick them.
	generator AnswerGenerator
	picker    AnswerPicker
	// keep track of past questions, use slices to maintain order.
	history map[string][]string
}

// NewEliza creates a new Eliza instance with teh given answer generator
// and answer picker.
func NewEliza(generator AnswerGenerator, picker AnswerPicker) *Eliza {
	eliza := Eliza{generator: generator, picker: picker, history: make(map[string][]string)}
	eliza.history["questions"] = []string{}
	eliza.history["answers"] = []string{}
	return &eliza
}

func (e *Eliza) saveQuestion(question string) {
	e.appendToList("questions", question)
}

func (e *Eliza) saveAnswer(answer string) {
	e.appendToList("answers", answer)
}

func (e *Eliza) appendToList(key, val string) {
	e.history[key] = append(e.history[key], val)
}

// GoAsk is the "main" exported function Eliza is needed for.
// it takes a single question in string format, and gives back
// a single response also in string format.
func (e *Eliza) GoAsk(question string) string {
	e.saveQuestion(question)
	answers := e.generator.GenerateAnswers(question)
	answer := e.picker.PickAnswer(answers)
	e.saveAnswer(answer)
	return answer
}

// Questions returns a list of all asked questions
func (e *Eliza) Questions() []string {
	return []string(e.history["questions"])
}

// Answers returns a list of all given answers.
func (e *Eliza) Answers() []string {
	return []string(e.history["answers"])
}

// https://github.com/golang/go/wiki/CodeReviewComments#interfaces
// The documentation states that interfaces belong in the package that is
// going to use the interface type, not in with the implementations.

// AnswerGenerator should be able to give back a slice of answers when given a question.
type AnswerGenerator interface {
	GenerateAnswers(question string) []string
}

// AnswerPicker is in charge of picking a response out of a list of answers.
type AnswerPicker interface {
	PickAnswer(answers []string) string
}
