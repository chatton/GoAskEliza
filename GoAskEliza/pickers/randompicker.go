package pickers

import (
    "math/rand"
    "time"
)

type RandomPicker struct {} // this implementation has no state.

func (picker RandomPicker) PickAnswer(answers []string) string{
    rand.Seed(time.Now().Unix()) // seed so we don't get the same value every time
    index := rand.Intn(len(answers)) // index between 0 -> No. answers
    return answers[index] // pick the random answer.
}