package pickers

import (
	"math/rand"
	"time"
)

//RandomPicker has no state, does not prioritse any
//answers over any others.
type RandomPicker struct{}

func (picker *RandomPicker) PickAnswer(answers []string) string {
	index := rand.Intn(len(answers)) // index between 0 -> No. answers
	answer := answers[index]
	return answer
}

func NewRandomPicker() *RandomPicker {
	rand.Seed(time.Now().Unix()) // seed so we don't get the same value every time
	return &RandomPicker{}
}
