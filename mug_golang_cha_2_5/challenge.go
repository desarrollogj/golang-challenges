package main

import (
	"fmt"
	"math"
)

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

type LastMinuteEmployee struct {
	Employee
	LastMinutePassenger
}

func (l LastMinuteEmployee) getDiscount() float64 {
	discount := l.Employee.getDiscount() + l.LastMinutePassenger.getDiscount()
	if discount > 1 {
		discount = 1
	}

	return discount
}

type FlightSeat interface {
	getCompleteName() string
	getDiscount() float64
}

func getFlightIncome(seats []FlightSeat, basePrice float64) float64 {
	var income float64

	for _, seat := range seats {
		income += getFlightSeatIncome(seat, basePrice)
	}

	return income
}

func getFlightSeatIncome(seat FlightSeat, basePrice float64) float64 {
	return basePrice - (basePrice * seat.getDiscount())
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

	var otherTom LastMinuteEmployee
	otherTom.firstName = "Tom"
	otherTom.lastName = "Cruise"

	basePrice := 180.0
	passengers := []FlightSeat{tom, denzel, leo, otherTom}

	for _, passenger := range passengers {
		price := getFlightSeatIncome(passenger, basePrice)
		payOrReceive := "pay"
		if price < 0 {
			payOrReceive = "receive"
			price = math.Abs(price)
		}

		fmt.Printf("%s will %s %.2f for the ticket\n", passenger.getCompleteName(), payOrReceive, price)
	}

	income := getFlightIncome(passengers, basePrice)
	fmt.Printf("Flight total income: %.2f\n", income)
}
