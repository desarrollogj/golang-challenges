package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	input := "Join the navy"
	output := "yvan eht nioJ"

	result := Reverse(input)
	if result != output {
		t.Errorf("Reverse(%s) == %s, want %s", input, result, output)
	}
}

func BenchmarkReverse(b *testing.B) {
	inputs := [...]string{"Uno", "Dos", "Tres", "Cuatro", "Cinco"}
	for _, input := range inputs {
		Reverse(input)
	}
}

func ExampleReverse() {
	fmt.Println(Reverse("Join the navy"))
	// Output: yvan eht nioJ
}
