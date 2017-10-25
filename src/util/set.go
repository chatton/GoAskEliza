package util

import (
	"errors"
	"math/rand"
	"time"
)

// I consulted this post on how to emulate a set data-structure in go
// https://softwareengineering.stackexchange.com/questions/177428/sets-data-structure-in-golang

// no need to support removal of elements currently. We only need
// to be able to Add and check for elements via Contains

type StringSet struct { // mimic a set using a map of string -> bool
	set map[string]bool
}

func (set *StringSet) Add(s string) {
	set.set[s] = true
}

func (set *StringSet) Contains(s string) bool {
	_, ok := set.set[s] // don't care about the value, just if it was there.
	return ok
}

func (set *StringSet) Size() int {
	return len(set.set)
}

func (set *StringSet) AsSlice() []string {
	allStrings := []string{}
	for key := range set.set {
		allStrings = append(allStrings, key)
	}
	return allStrings
}

func (set *StringSet) RandomValue() (string, error) {
	if set.IsEmpty() {
		return "", errors.New("Set is empty")
	}

	rand.Seed(time.Now().UTC().UnixNano())
	values := set.AsSlice()
	return values[rand.Intn(len(values))], nil
}

func (set *StringSet) IsEmpty() bool {
	return set.Size() == 0
}

func NewStringSet() *StringSet {
	return &StringSet{make(map[string]bool)}
}
