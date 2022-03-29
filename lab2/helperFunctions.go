package lab2

import (
	"math"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func makeByte(bitstring string) byte {
	if len(bitstring) < 8 {
		bitstring = bitstring + "1"
	}

	for {
		if len(bitstring) < 8 {
			bitstring = bitstring + "0"
		} else {
			break
		}
	}

	val := 0

	for i := 0; i < 8; i++ {
		bit := bitstring[i]
		if bit == '1' {
			val += int(math.Pow(2.0, float64(7.0-i)))
		}
	}

	return byte(val)
}

func makeBitstring(b byte) string {
	bitstring := strconv.FormatUint(uint64(b), 2)
	l := len(bitstring)

	for i := 0; i < 8-l; i++ {
		bitstring = "0" + bitstring
	}

	return bitstring
}

func getTagValue(bitstring string) float64 {
	tag := 0.0

	for i := 0; i < len(bitstring); i++ {
		if bitstring[i] == '1' {
			tag += math.Pow(0.5, float64(i+1))
		}
	}

	return tag
}
