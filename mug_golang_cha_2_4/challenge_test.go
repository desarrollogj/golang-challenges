package main

import (
	"testing"
)

func TestGetFlightIncomeWithAllPassengerTypes(t *testing.T) {
	inputSeats := []FlightSeat{Passenger{}, LastMinutePassenger{}, Employee{}}
	inputBasePrice := 150.0
	output := 225.00

	result := getFlightIncome(inputSeats, inputBasePrice)
	if result != output {
		t.Errorf("getFlightIncome(%+v,%.2f) == %.2f, want %.2f", inputSeats, inputBasePrice, result, output)
	}
}

func TestGetFlightIncomeWithPassengers(t *testing.T) {
	inputSeats := []FlightSeat{Passenger{}, Passenger{}, Passenger{}}
	inputBasePrice := 150.0
	output := 450.0

	result := getFlightIncome(inputSeats, inputBasePrice)
	if result != output {
		t.Errorf("getFlightIncome(%+v,%.2f) == %.2f, want %.2f", inputSeats, inputBasePrice, result, output)
	}
}

func TestGetFlightIncomeWithLastMinutePassengers(t *testing.T) {
	inputSeats := []FlightSeat{LastMinutePassenger{}, LastMinutePassenger{}, LastMinutePassenger{}}
	inputBasePrice := 150.0
	output := 225.0

	result := getFlightIncome(inputSeats, inputBasePrice)
	if result != output {
		t.Errorf("getFlightIncome(%+v,%.2f) == %.2f, want %.2f", inputSeats, inputBasePrice, result, output)
	}
}

func TestGetFlightIncomeWithEmployees(t *testing.T) {
	inputSeats := []FlightSeat{Employee{}, Employee{}, Employee{}}
	inputBasePrice := 150.0
	output := 0.0

	result := getFlightIncome(inputSeats, inputBasePrice)
	if result != output {
		t.Errorf("getFlightIncome(%+v,%.2f) == %.2f, want %.2f", inputSeats, inputBasePrice, result, output)
	}
}
