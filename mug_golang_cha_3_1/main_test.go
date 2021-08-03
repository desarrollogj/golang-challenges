package main

import "testing"

func TestSquaresPlusCubes(t *testing.T) {
	input := 6
	output := 252

	result := SquaresPlusCubes(input)
	if result != output {
		t.Errorf("SquaresPlusCubes(%d) == %d, want %d", input, result, output)
	}
}

func TestSquaresPlusCubesChan(t *testing.T) {
	input := 6
	output := 252

	result := SquaresPlusCubesChan(input)
	if result != output {
		t.Errorf("SquaresPlusCubesChan(%d) == %d, want %d", input, result, output)
	}
}

func BenchmarkSquaresPlusCubes(b *testing.B) {
	input := 1
	for input < 100000 {
		SquaresPlusCubes(input)
		input++
	}
}

func BenchmarkSquaresPlusCubesChan(b *testing.B) {
	input := 1
	for input < 100000 {
		SquaresPlusCubesChan(input)
		input++
	}
}
