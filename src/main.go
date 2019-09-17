package main

import (
	"GoAskEliza/src/eliza"
	"GoAskEliza/src/generators"
	"GoAskEliza/src/pickers"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var el *eliza.Eliza

type History struct {
	Answers, Questions []string
}

func main() {

	connectionStr, ok := os.LookupEnv("CONNECTION_STRING")
	if !ok {
		fmt.Println("no CONNECTION_STRING provided")
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.TODO(), 2*time.Second)
	// "mongodb://localhost:20000"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionStr))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gen := generators.NewRegexGenerator("./data/pattern-responses.dat")
	picker := pickers.NewPrefersNewPicker()
	el = eliza.NewEliza(gen, picker, client)
	http.HandleFunc("/ask", handleAsk)
	http.HandleFunc("/history", handleHistory)
	http.Handle("/", http.FileServer(http.Dir("./web")))
	fmt.Println("Starting web server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

func handleAsk(w http.ResponseWriter, r *http.Request) {
	if userHasQuestion(r) {
		var userQuestion string
		switch r.Method { // can handle both GET and POST requests.
		case "GET":
			userQuestion = r.URL.Query().Get("question") // extract GET parameter
		case "POST":
			userQuestion = r.FormValue("question") // extracts POST parameter
		}

		answer, err := el.GoAsk(userQuestion) // passes the user question to the Eliza struct to get an answer for the question.
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, answer) // write the answer back.
	}
}

func handleHistory(w http.ResponseWriter, r *http.Request) {
	answers, err := el.Answers()
	if err != nil {
		panic(err)
	}
	questions, err := el.Questions()
	if err != nil {
		panic(err)
	}
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
