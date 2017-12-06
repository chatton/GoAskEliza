package main

import (
	"fmt"
	"net/http"
	"strings"

	"./eliza"
	"./generators"
	"./pickers"
)

// hold onto one instance of eliza per user.
var elizas map[string]*eliza.Eliza

type History struct {
	Answers, Questions []string
}

func newEliza() *eliza.Eliza {
	gen := generators.NewRegexGenerator("./data/pattern-responses.dat")
	picker := pickers.NewPrefersNewPicker()
	return eliza.NewEliza(gen, picker)
}

func main() {
	elizas = make(map[string]*eliza.Eliza)
	http.HandleFunc("/ask", handleAsk)
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.ListenAndServe(":8080", nil)
}

func handleAsk(w http.ResponseWriter, r *http.Request) {
	if userHasQuestion(r) {
		
		userId := r.FormValue("id")
		if _, ok := elizas[userId]; !ok {
			elizas[userId] = newEliza()
		}

		usersEliza , _ := elizas[userId]

		userQuestion := r.FormValue("question") // extracts POST parameter
		fmt.Println(fmt.Sprintf("User Id[%s] Asked: %s", userId, userQuestion))
		answer := usersEliza.GoAsk(userQuestion) // passes the user question to the Eliza struct to get an answer for the question.
		fmt.Fprintf(w, answer)           // write the answer back.	
		fmt.Println(fmt.Sprintf("Eliza: %s", answer))
	}
}

func userHasQuestion(r *http.Request) bool {
	// the user has a question if they have a non-empty question.
	return strings.TrimSpace(r.FormValue("question")) != "" || strings.TrimSpace(r.URL.Query().Get("question")) != ""
}
