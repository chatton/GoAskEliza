package pickers

import (
	"../util"
	"math/rand"
	"time"
)

type RandomPicker struct {
	// keep track of all the questions already picked.
	// this way we can prioritize the other questions, to seem more "real"
	pickedAnswers *util.StringSet
}

func (picker RandomPicker) PickAnswer(answers []string) string {
	rand.Seed(time.Now().Unix())     // seed so we don't get the same value every time
	index := rand.Intn(len(answers)) // index between 0 -> No. answers
	answer := answers[index]
	if picker.pickedAnswers.Contains(answer) { // already picked this answer
		index = rand.Intn(len(answers)) // pick another one
		answer = answers[index]         // can still be the same number, this just reduces
		// the likeyhood of a same answer being picked.
	}
	return answer
}

func NewRandomPicker() *RandomPicker {
	return &RandomPicker{pickedAnswers: util.NewStringSet()}
}
