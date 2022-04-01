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

func CalculateEntropy(d *dictionary) float64 {
	entropy := 0.0

	for _, s := range d.symbols {
		entropy += float64(s.count) / float64(d.totalCount) * -math.Log2(float64(s.count)/float64(d.totalCount))
	}

	return entropy
}

func CalculateCompressionRate(inputSize uint64, outputSize uint64) float64 {
	compressionRate := float64(inputSize) / float64(outputSize)

	return compressionRate
}

func CalculateAverageCodeLength(inputSize uint64, outputSize uint64) float64 {
	avgCodeLength := 8 / CalculateCompressionRate(inputSize, outputSize)

	return avgCodeLength

}
