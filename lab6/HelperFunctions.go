package lab6

import (
	"math"
	"strconv"
)

func byteToString(b byte, bitrate int) string {
	s := strconv.FormatUint(uint64(b), 2)

	length := len(s)
	for i := 0; i < bitrate-length; i++ {
		s = "0" + s
	}

	return s
}

func stringToByte(s string) byte {
	var b byte = 0
	length := len(s)

	for i := 0; i < length; i++ {
		if s[i] == '1' {
			b += byte(math.Pow(2.0, float64(length-i-1)))
		}
	}

	return b
}

func abs(a int, b int) int {
	return int(math.Abs(float64(a) - float64(b)))
}

func indexof(arr []int, elem int) int {
	for i, v := range arr {
		if elem == v {
			return i
		}
	}

	return -1
}
