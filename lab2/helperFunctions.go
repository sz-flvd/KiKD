package lab2

import (
	"math"
	"strconv"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func MakeByte(bitstring string) byte {
	for {
		if len(bitstring) < 8 {
			bitstring = bitstring + "0"
		} else {
			break
		}
	}

	val := 0

	for i := 0; i < 8; i++ {
		if bitstring[i] == '1' {
			val += int(math.Pow(2.0, float64(7-i)))
		}
	}

	return byte(val)
}

func MakeBitstring(b byte) string {
	bitstring := strconv.FormatUint(uint64(b), 2)
	l := len(bitstring)

	for i := 0; i < 8-l; i++ {
		bitstring = "0" + bitstring
	}

	return bitstring
}

func FloatToBitstring(f float64, precision int) string {
	bitstring := ""
	approx := 0.0
	zeroCounter := 0

	for i := 1; i <= precision; {
		if f >= approx+math.Pow(0.5, float64(i)) {
			for j := 0; j < zeroCounter; j++ {
				bitstring = bitstring + "0"
			}
			zeroCounter = 0
			bitstring = bitstring + "1"
			approx += math.Pow(0.5, float64(i))
			i++
		} else {
			zeroCounter++
			i++
		}
	}

	return bitstring
}

func GetTagValue(bitstring string) float64 {
	tag := 0.0

	for i := 0; i < len(bitstring); i++ {
		if bitstring[i] == '1' {
			tag += math.Pow(0.5, float64(i+1))
		}
	}

	return tag
}
