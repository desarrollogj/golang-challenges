package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(Reverse("Anita LAVA la TINA"))
}

func Reverse(input string) string {
	output := ""

	for i := len(input) - 1; i >= 0; i = i - 1 {
		output += string(reverseIfIsVowel(rune(input[i])))
	}

	return output
}

func reverseIfIsVowel(r rune) rune {
	if isVowel(r) {
		return reverseCasing(r)
	}

	return r
}

func isVowel(r rune) bool {
	r = unicode.ToLower(r)
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	for _, v := range vowels {
		if v == r {
			return true
		}
	}

	return false
}

func reverseCasing(r rune) rune {
	if unicode.IsLower(r) {
		return unicode.ToUpper(r)
	}

	return unicode.ToLower(r)
}
