package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func (p Person) getCompleteName() string {
	return p.firstName + " " + p.lastName
}

type Passenger struct {
	Person
}

func (p Passenger) getDiscount() float64 {
	return 0
}

type LastMinutePassenger struct {
	Passenger
}

func (l LastMinutePassenger) getDiscount() float64 {
	return 0.5
}

type Employee struct {
	Person
}

func (e Employee) getDiscount() float64 {
	return 1
}

type FlightSeat interface {
	getCompleteName() string
	getDiscount() float64
}

func getFlightIncome(seats []FlightSeat, basePrice float64) float64 {
	var income float64

	for _, seat := range seats {
		income += basePrice - (basePrice * seat.getDiscount())
	}

	return income
}

func main() {
	var tom Passenger
	tom.firstName = "Tom"
	tom.lastName = "Hanks"

	var denzel LastMinutePassenger
	denzel.firstName = "Denzel"
	denzel.lastName = "Washington"

	var leo Employee
	leo.firstName = "Leonardo"
	leo.lastName = "DiCaprio"

	basePrice := 180.0
	passengers := []FlightSeat{tom, denzel, leo}

	// For check struct methods
	/*for _, passenger := range passengers {
		price := basePrice - (basePrice * passenger.getDiscount())
		fmt.Printf("%s will pay %.2f for the ticket\n", passenger.getCompleteName(), price)
	}*/

	income := getFlightIncome(passengers, basePrice)
	fmt.Printf("Flight total income: %.2f\n", income)
}
