package util

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

func NewStringSet() *StringSet {
    return &StringSet{make(map[string]bool)}
}