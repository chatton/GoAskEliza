package pickers

import (
	"GoAskEliza/src/util"
	"math/rand"
	"time"

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
	// I had the idea to shuffle the answer list and pick from the start
	// I found this post https://stackoverflow.com/questions/12264789/shuffle-array-in-go
	// and adapted it to come to this solution.

	indices := rand.Perm(len(answers)) // get all possible indices of the answer slice. e.g. [0,2,1]
	var answer string
	// here index IS the "value"
	for _, index := range indices {
		// pick out that answer
		answer = answers[index]
		// an answer we haven't seen before.
		if !picker.pickedAnswers.Contains(answer) {
			break
		}
	}
	// remember the answer
	picker.pickedAnswers.Add(answer)
	return answer
}

func NewPrefersNewPicker() *PrefersNewPicker {
	rand.Seed(time.Now().Unix()) // seed so we don't get the same value every time
	return &PrefersNewPicker{pickedAnswers: util.NewStringSet()}
}
