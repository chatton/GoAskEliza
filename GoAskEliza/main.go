package main

import (
    "./eliza"
    "fmt"
)

func main() {
    eliza := eliza.New()
    fmt.Println(eliza.GoAsk("How are you?"))
}