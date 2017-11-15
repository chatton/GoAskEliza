package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"./eliza"
	"./generators"
	"./pickers"
)

var el *eliza.Eliza

type History struct {
	Answers, Questions []string
}

func main() {
	gen := generators.NewRegexGenerator("./data/pattern-responses.dat")
	picker := pickers.NewPrefersNewPicker()
	el = eliza.NewEliza(gen, picker)

	http.HandleFunc("/ask", handleAsk)
	http.HandleFunc("/history", handleHistory)
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.ListenAndServe(":8080", nil)
}

func handleAsk(w http.ResponseWriter, r *http.Request) {
	if userHasQuestion(r) {
		var userQuestion string
		switch r.Method {
		case "GET":
			userQuestion = r.URL.Query().Get("question") // extract GET parameter
		case "POST":
			userQuestion = r.FormValue("question") // extracts POST parameter
		}

		answer := el.GoAsk(userQuestion) // passes the user question to the Eliza struct to get an answer for the question.
		fmt.Fprintf(w, answer)           // write the answer back.
	}
}

func handleHistory(w http.ResponseWriter, r *http.Request) {
	answers := el.Answers()
	questions := el.Questions()

	// I consulted this article on how to use json.Marshal correctly
	// https://golangcode.com/json-encode-an-array-of-objects/
	history := History{answers, questions}
	historyJSON, err := json.Marshal(history)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "%s", historyJSON)
}

func userHasQuestion(r *http.Request) bool {
	// the user has a question if they have a non-empty question.
	return strings.TrimSpace(r.FormValue("question")) != "" || strings.TrimSpace(r.URL.Query().Get("question")) != ""
}
