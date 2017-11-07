package eliza

import (
	"net/http"
	"fmt"
	"strings"
)

type server struct{
	el *Eliza
}

func NewServer(el *Eliza) *server {
	server := &server{}
	server.el = el
	return server
}

func (server *server) Start(){
	http.HandleFunc("/ask", server.handleAsk)
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.ListenAndServe(":8080", nil)
}
/*
func serveIndex(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r, "./web/index.html")
}
*/
func (server *server) handleAsk(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if userHasInput(r) { 
		userQuestion := r.FormValue("question") // the value gets passed in in the input-form.
		fmt.Println(userQuestion)
		answer := server.el.GoAsk(userQuestion) // passes the user question to the Eliza struct to get an answer for the question.
		fmt.Println(answer)
		fmt.Fprintf(w, answer)
	}
	//serveIndex(w,r)
}

func userHasInput(r *http.Request) bool {
	// the user has a question if they have a non-empty question.
	return strings.TrimSpace(r.FormValue("question")) != ""
}