package main

import (
	"TSPs_Hopfield/algorithm"
	"fmt"
)


func main() {
	cities := []algorithm.City{
		{Name: "A", X: 20, Y: 79},
		{Name: "B", X: 61, Y: 72},
		{Name: "C", X: 78, Y: 82},
		{Name: "D", X: 23, Y: 16},
		{Name: "E", X: 15, Y: 24},
		{Name: "F", X: 87, Y: 9},
		{Name: "G", X: 98, Y: 37},
		{Name: "H", X: 45, Y: 98},
	}

	tsp := algorithm.NewTSP(cities)

	// Sample configuration (tour)
	states := [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 1},
	}
	// Set the values for A, B, C, and D (You need to tune these values based on your problem)
	A := 1.0
	B := 1.0
	C := 1.0
	D := 1.0

	// Calculate the energy using the HopfieldEnergy function
	energy := tsp.HopfieldEnergy(states, A, B, C, D)

	// Print the calculated energy
	fmt.Printf("Energy: %f\n", energy)

}

