package eliza

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type server struct {
	el *Eliza
}

type History struct {
	Answers, Questions []string
}

func NewServer(el *Eliza) *server {
	server := &server{}
	server.el = el
	return server
}

func (server *server) Start() {
	http.HandleFunc("/ask", server.handleAsk)
	http.HandleFunc("/history", server.handleHistory)
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.ListenAndServe(":8080", nil)
}

func (server *server) handleAsk(w http.ResponseWriter, r *http.Request) {
	if userHasQuestion(r) {
		userQuestion := r.FormValue("question") // the value gets passed in in the input-form.
		answer := server.el.GoAsk(userQuestion) // passes the user question to the Eliza struct to get an answer for the question.
		fmt.Fprintf(w, answer)                  // write the answer back.
	}
}

func (server *server) handleHistory(w http.ResponseWriter, r *http.Request) {
	answers := server.el.Answers()
	questions := server.el.Questions()

	// I consulted this article on how to use json.Marshal correctly
	// https://golangcode.com/json-encode-an-array-of-objects/
	history := History{answers, questions}
	historyJson, err := json.Marshal(history)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "%s", historyJson)
}

func userHasQuestion(r *http.Request) bool {
	// the user has a question if they have a non-empty question.
	return strings.TrimSpace(r.FormValue("question")) != ""
}
