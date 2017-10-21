package main

import (
	"./eliza"
	"./generators"
	"./pickers"
	"fmt"
	"net/http"
)

var e *eliza.Eliza

func main() {
	var g eliza.AnswerGenerator // interface
	// implementation
	g = generators.NewRegexGenerator("./data/pattern-responses.dat")

	var p eliza.AnswerPicker
	p = pickers.NewRandomPicker()

	e = eliza.NewEliza(g, p)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// serves up index.html when a request is made to "/"
		http.ServeFile(w, r, "./html/index.html")
	})

	// when a request is made to /ask, the askEliza function will execute.
	http.HandleFunc("/ask", askEliza)

	http.ListenAndServe(":9999", nil)
}

func hasQuesion(r *http.Request) bool {
	return getQuestion(r) != ""
}

func getQuestion(r *http.Request) string {
	return r.URL.Query().Get("question")
}

func askEliza(w http.ResponseWriter, r *http.Request) {
	var question string
	if hasQuesion(r) {
		question = getQuestion(r)
		response := e.GoAsk(question)
		fmt.Println(response)
	}
}
