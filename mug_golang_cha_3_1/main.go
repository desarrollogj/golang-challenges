package main

import "fmt"

func Squares(number int) (sum int) {
	sum = 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	return
}

func SquaresChan(number int, c chan int) {
	c <- Squares(number)
}

func Cubes(number int) (sum int) {
	sum = 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	return
}

func CubesChan(number int, c chan int) {
	c <- Cubes(number)
}

func SquaresPlusCubes(number int) int {
	return Squares(number) + Cubes(number)
}

func SquaresPlusCubesChan(number int) int {
	c := make(chan int)
	go SquaresChan(number, c)
	go CubesChan(number, c)
	result := <-c
	result += <-c

	return result
}

func main() {
	fmt.Printf("SquaresPlusCubes of 6: %d\n", SquaresPlusCubes(6))
	fmt.Printf("SquaresPlusCubesChan of 6: %d\n", SquaresPlusCubesChan(6))
}
