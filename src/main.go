package main

import (
	"./eliza"
	"./generators"
	"./pickers"
	"fmt"
	"net/http"
)

func main() {
	var g eliza.AnswerGenerator // interface
	// implementation
	g = generators.NewRegexGenerator("./data/pattern-responses.dat")

	var p eliza.AnswerPicker
	p = pickers.NewRandomPicker()

	e := eliza.NewEliza(g, p)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "./html/index.html")
	})

	http.ListenAndServe(":9999", nil);
	fmt.Println(e.GoAsk("I like waffles %s!"))
}
