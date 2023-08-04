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
		{1, 0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 1},
	}
	// Set the values for A, B, C, and D
	A := 0.1
	B := 0.1
	C := 0.0000001
	D := 0.2

	// Calculate the energy using the HopfieldEnergy function
	energy := tsp.HopfieldEnergy(states, A, B, C, D)

	// Print the calculated energy
	fmt.Printf("Initial Energy: %f\n", energy)

	weights := tsp.GenerateSymmetricWeightMatrix(A, B, C, D)

	// usages.PrintMatrix(weights)



	// Set the convergence threshold
	convergenceThreshold := 0.0001

	// Run the HopfieldDynamic function to update states until convergence
	tsp.HopfieldDynamic(states,weights, A, B, C, D, convergenceThreshold)

	// // Decode the solution to get the tour order
	// finalTourOrder := tsp.DecodeSolution(states)

	// // Calculate the total tour length
	// totalTourLength := tsp.CalculateTotalTourLength(finalTourOrder)

	// fmt.Println("Final Tour Order:")
	// for i, city := range finalTourOrder {
	// 	fmt.Printf("%d. %s\n", i+1, city.Name)
	// }

	// fmt.Printf("Total Tour Length: %.2f\n", totalTourLength)

}

