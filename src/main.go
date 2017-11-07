package main

import (
	"./eliza"
	"./generators"
	"./pickers"
)

func main(){
	gen := generators.NewRegexGenerator("./data/pattern-responses.dat")
	picker := pickers.NewPrefersNewPicker()
	el := eliza.NewEliza(gen, picker)
	server := eliza.NewServer(el)
	server.Start()
}