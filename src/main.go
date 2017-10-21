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

func hasQuestion(r *http.Request) bool {
	return getQuestion(r) != ""
}

// helper function to get the question from the url.
func getQuestion(r *http.Request) string {
	return r.URL.Query().Get("question")
}

// can tell if the user has been here before if they have cookies from us.
func usersFirstTime(r *http.Request) bool {
	return len(r.Cookies()) == 0
}

func askEliza(w http.ResponseWriter, r *http.Request) {
	var question string
	if hasQuestion(r) {
		question = getQuestion(r)
		response := e.GoAsk(question)
		fmt.Println(response)
	}
}
