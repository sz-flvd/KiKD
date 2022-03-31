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

func TagToBitstring(tag uint64, precision int) string {
	bitstring := strconv.FormatUint(tag, 2)
	l := len(bitstring)

	for i := 0; i < precision-l; i++ {
		bitstring = "0" + bitstring
	}

	return bitstring
}

func GetTagValue(bitstring string, precision uint8) uint64 {
	tag := uint64(0)

	for i := 0; i < len(bitstring); i++ {
		if bitstring[i] == '1' {
			tag += uint64(math.Pow(2, float64(int(precision)-i-1)))
		}
	}

	return tag
}
