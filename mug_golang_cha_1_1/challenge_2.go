package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(Reverse("Join the Navy"))
	fmt.Println(Reverse2("Join the Navy"))
}

func Reverse(input string) string {
	output := ""
	for i := len(input) - 1; i >= 0; i = i - 1 {
		output += string(input[i])
	}
	return output
}

func Reverse2(s string) string {
	reversed := make([]rune, utf8.RuneCountInString(s))
	topIndex := len(s) - 1

	for i, r := range s {
		reversed[topIndex-i] = r
	}

	return string(reversed)
}
