package util

import (
	"errors"
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

func (set *StringSet) Remove(s string) error {
	if !set.Contains(s) {
		return errors.New("Element: " + s + " does not exist in the set.")
	}
	delete(set.set, s) // delete instead of setting to false, potential memory leak keeping values that "aren't there"
	return nil         // no error, the element was removed successfully
}

func (set *StringSet) Size() int {
	return len(set.set)
}

func (set *StringSet) Values() []string {
	allStrings := make([]string, len(set.set))
	index := 0
	for key := range set.set {
		allStrings[index] = key
		index++
	}
	return allStrings
}

func (set *StringSet) IsEmpty() bool {
	return set.Size() == 0
}

func NewStringSet(initialVals ...string) *StringSet {

	set := &StringSet{make(map[string]bool)}
	for _, val := range initialVals {
		set.Add(val)
	}
	return set
}
