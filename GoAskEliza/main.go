package main

import (
    "./eliza"
    "./eliza/answers/generator"
    "./eliza/answers/picker"
    "fmt"
)

func main() {
    var g generator.AnswerGenerator // interface
    g = generator.RegexGenerator{} // implementation

    var p picker.RandomPicker
    p = picker.RandomPicker{}

    e := eliza.NewEliza(g, p)
    fmt.Println(e.GoAsk("How are you?"))
}