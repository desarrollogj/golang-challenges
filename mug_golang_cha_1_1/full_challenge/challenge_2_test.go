package main

import (
	"testing"
)

type data struct {
	InA      string
	Expected string
}

var listOfData = []data{
	{InA: "Join the Navy 1", Expected: ""},
	{InA: "Join the Navy 2", Expected: ""},
	{InA: "Join the Navy 3", Expected: ""},
	{InA: "Join the Navy 4", Expected: ""},
	{InA: "Join the Navy 5", Expected: ""},
}

func BenchmarkReverse_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range listOfData {
			Reverse2(value.InA)
		}
	}
}

func BenchmarkReverseVariant_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range listOfData {
			ReverseVariant2(value.InA)
		}
	}
}
