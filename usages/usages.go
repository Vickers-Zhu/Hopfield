package usages

import "fmt"

func PrintMatrix(matrix [][]float64) {
	fmt.Print("     ")
	for i := range matrix {
		fmt.Printf("City %d   ", i+1)
	}
	fmt.Println()

	for i, row := range matrix {
		fmt.Printf("City %d ", i+1)
		for _, value := range row {
			fmt.Printf("%.2f    ", value)
		}
		fmt.Println()
	}
}

func SumAllStates(states [][] int) float64 {
	sum := 0
	for i := 0; i < len(states); i++ {
		for j := 0; j < len(states[i]); j++ {
			sum += states[i][j]
		}
	}
	return float64(sum)
}
