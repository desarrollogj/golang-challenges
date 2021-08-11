/*Package main provides Reverses functions.
-------------------------------[ go test -cpu=4 -bench=. ]--------------------------------
Run that command to validate the performance of each function.
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(Reverse2("Join the Navy"))
	fmt.Println(ReverseVariant2("Join the Navy"))
}

func Reverse2(input string) string {
	output := ""
	for i := len(input) - 1; i >= 0; i = i - 1 {
		output += string(input[i])
	}
	return output
}

func ReverseVariant2(s string) string {
	reversed := make([]rune, utf8.RuneCountInString(s))
	topIndex := len(s) - 1

	for i, r := range s {
		reversed[topIndex-i] = r
	}

	return string(reversed)
}
