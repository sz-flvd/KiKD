package lab3

import "math"

func first(b []byte, _ error) []byte {
	return b
}

func entropy(data map[byte]uint) float64 {
	entropy := 0.0
	sum := uint(0)

	for _, s := range data {
		sum += s
	}

	for _, s := range data {
		entropy += float64(s) / float64(sum) * -math.Log2(float64(s)/float64(sum))
	}

	return entropy
}
