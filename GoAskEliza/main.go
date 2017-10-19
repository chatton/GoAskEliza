package main

import (
    "./eliza"
    "./generators"
    "./pickers"
    "fmt"
)

func main() {
    var g eliza.AnswerGenerator // interface
    // implementation
    g = generators.NewRegexGenerator("./data/pattern-responses.dat") 

    var p eliza.AnswerPicker
    p = pickers.NewRandomPicker()

    e := eliza.NewEliza(g, p)
    fmt.Println(e.GoAsk("I like waffles!"))
}