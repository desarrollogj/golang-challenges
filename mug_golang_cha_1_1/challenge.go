package main

import (
	"fmt"
)

func main() {
	fmt.Println(Reverse("Join the Navy"))
}

func Reverse(input string) string {
	output := ""
	inputRune := []rune(input)
	for i := len(input) - 1; i >= 0; i = i - 1 {
		output += string(inputRune[i])
	}
	return output
}
