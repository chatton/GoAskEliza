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

func NewEliza(generator AnswerGenerator, picker AnswerPicker) *Eliza {
	eliza := Eliza{generator: generator, picker: picker, history: make(map[string][]string)}
	eliza.history["questions"] = []string{}
	eliza.history["answers"] = []string{}
	return &eliza
}

func (e *Eliza) saveQuestion(question string) {
    e.history["questions"] = append(e.history["questions"], question)
}

func (e *Eliza) saveAnswer(answer string) {
    e.history["answers"] = append(e.history["answers"], answer)
}

func (e *Eliza) GoAsk(question string) string {
    e.saveQuestion(question)
	answers := e.generator.GenerateAnswers(question)
    answer := e.picker.PickAnswer(answers)
    e.saveAnswer(answer)
	return answer
}

func (e *Eliza) Questions() []string {
	return []string(e.history["questions"])
}

func (e *Eliza) Answers() []string {
	return []string(e.history["answers"])
}

func (e *Eliza) Greet(firstTime bool) string {
	if firstTime {
		return "Hi, my name is Eliza, it's nice to meet you."
	}
	return "Welcome back. Let's continue."
}
 
// https://github.com/golang/go/wiki/CodeReviewComments#interfaces
// The documentation states that interfaces belong in the package that is
// going to use the interface type, not in with the implementations.

// an answer generator should be able to give back a slice of answers when given a question.
type AnswerGenerator interface {
	GenerateAnswers(question string) []string
}

// answer picker is in charge of picking a response out of a list of answers.
type AnswerPicker interface {
	PickAnswer(answers []string) string
}
