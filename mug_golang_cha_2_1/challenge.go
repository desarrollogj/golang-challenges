package main

import (
	"flag"
	"fmt"
	"unicode/utf8"
)

func main() {
	textPtr := flag.String("text", "Example", "a string")
	flag.Parse()

	fmt.Println(Reverse(*textPtr))
}

func Reverse(s string) string {
	reversed := make([]rune, utf8.RuneCountInString(s))
	topIndex := len(s) - 1

	for i, r := range s {
		reversed[topIndex-i] = r
	}

	return string(reversed)
}
