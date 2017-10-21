package pickers

import (
    "math/rand"
    "time"
)

// I consulted this post on how to emulate a set data-strucure in go
// https://softwareengineering.stackexchange.com/questions/177428/sets-data-structure-in-golang

type StringSet struct { // mimic a set using a map of string -> bool
    set map[string]bool
}

func (set *StringSet) Add(s string)  {
    set.set[s] = true
}

func (set *StringSet) Contains(s string)  bool {
    _, ok := set.set[s] // don't care about the value, just if it was there.
    return ok
}

func NewStringSet() *StringSet {
    return &StringSet{make(map[string]bool)}
}

type RandomPicker struct {
    // keep track of all the questions already picked.
    // this way we can prioritize the other questions, to seem more "real"
    pickedAnswers *StringSet
} 

func (picker RandomPicker) PickAnswer(answers []string) string{
    rand.Seed(time.Now().Unix()) // seed so we don't get the same value every time
    index := rand.Intn(len(answers)) // index between 0 -> No. answers
    answer := answers[index]
    if picker.pickedAnswers.Contains(answer) { // already picked this answer
        index = rand.Intn(len(answers)) // pick another one
        answer = answers[index] // can still be the same number, this just reduces
        // the likeyhood of a same answer being picked.
    }
    return answer 
}

func NewRandomPicker() *RandomPicker {
    return &RandomPicker{pickedAnswers:NewStringSet()}
}