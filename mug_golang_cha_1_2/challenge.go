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
	inputRune := []rune(input)
	for i := len(input) - 1; i >= 0; i = i - 1 {
		char := inputRune[i]
		if IsVowel(char) {
			char = ChangeVowelCase(char)
		}
		output += string(char)
	}
	return output
}

func IsVowel(char rune) bool {
	vowels := [...]rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, vowel := range vowels {
		if char == vowel {
			return true
		}
	}
	return false
}

func ChangeVowelCase(vowel rune) rune {
	var output rune
	switch vowel {
	case 'a', 'e', 'i', 'o', 'u':
		output = unicode.ToUpper(vowel)
	case 'A', 'E', 'I', 'O', 'U':
		output = unicode.ToLower(vowel)
	default:
		output = vowel
	}
	return output
}
