package main

import (
	"KiKD/lab4"
	"fmt"
	"os"
)

func lab4Main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go file.tga")
		return
	}

	bytes := lab4.LoadTGAFile(os.Args[1])
	var predicted [][]int
	var predictedBytes [][]byte
	var entropy float64
	var entropyRed float64
	var entropyGreen float64
	var entropyBlue float64
	var best float64
	var bestRed float64
	var bestGreen float64
	var bestBlue float64
	bestPredicates := []int{0, 0, 0, 0}

	entropy, entropyRed, entropyGreen, entropyBlue = lab4.Entropy(&bytes)
	best, bestRed, bestGreen, bestBlue = entropy, entropyRed, entropyGreen, entropyBlue
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Predicate\tEntropy\t\tEntropy - R\tEntropy - G\tEntropy - B\n")
	fmt.Printf("Original file\t%f\t%f\t%f\t%f\n", entropy, entropyRed, entropyGreen, entropyBlue)

	for i := 1; i <= 8; i++ {
		predicted = lab4.Predict(&bytes, i)
		predictedBytes = lab4.IntArrayToByteArray(&predicted)
		entropy, entropyRed, entropyGreen, entropyBlue = lab4.Entropy(&predictedBytes)
		if entropy < best {
			best = entropy
			bestPredicates[0] = i
		}
		if entropyRed < bestRed {
			bestRed = entropyRed
			bestPredicates[1] = i
		}
		if entropyGreen < bestGreen {
			bestGreen = entropyGreen
			bestPredicates[2] = i
		}
		if entropyBlue < bestBlue {
			bestBlue = entropyBlue
			bestPredicates[3] = i
		}
		fmt.Printf("Predicate %d\t%f\t%f\t%f\t%f\n", i, entropy, entropyRed, entropyGreen, entropyBlue)
	}
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Best predicates:\t%d\t\t%d\t\t%d\t\t%d\n", bestPredicates[0], bestPredicates[1], bestPredicates[2], bestPredicates[3])
}

func main() {
	lab4Main()
}
