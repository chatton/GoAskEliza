package main

import "fmt"
import "./eliza"
import "./generators"
import "./pickers"
import "os"


/*
short tool to test the functionality of
Eliza. Can pass in a question and get back a single response.
*/

func main(){
	var g eliza.AnswerGenerator
	g = generators.NewRegexGenerator("./data/pattern-responses.dat")

	var p eliza.AnswerPicker
	p = pickers.NewRandomPicker()

	e := eliza.NewEliza(g, p)
	if len(os.Args) < 2{
		fmt.Println("usage: \"go run ask.go <question>\"")
		os.Exit(0)
	}
	question := os.Args[1]
	fmt.Println(e.GoAsk(question))

}