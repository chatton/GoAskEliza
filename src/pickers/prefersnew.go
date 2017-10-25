package pickers

import (
	"math/rand"
	"time"

	"../util"
)

//PrefersNewPicker will pick new answers if they are available.
//It will only pick duplicates if every possible answer
//has been seen already.
type PrefersNewPicker struct {
	// keep track of all the questions already picked.
	// this way we can prioritize the other questions, to seem more "real"
	pickedAnswers *util.StringSet
}

func (picker *PrefersNewPicker) PickAnswer(answers []string) string {
	index := rand.Intn(len(answers)) // index between 0 -> No. answers
	answer := answers[index]

	numSamePicks := 0
	for picker.pickedAnswers.Contains(answer) || numSamePicks == len(answers) { // already picked this answer
		numSamePicks++
		index = rand.Intn(len(answers)) // pick another one
		answer = answers[index]         // can still be the same number, this just reduces
		// the likeyhood of a same answer being picked.
	}

	picker.pickedAnswers.Add(answer)
	return answer
}

func NewPrefersNewPicker() *PrefersNewPicker {
	rand.Seed(time.Now().Unix()) // seed so we don't get the same value every time
	return &PrefersNewPicker{pickedAnswers: util.NewStringSet()}
}
