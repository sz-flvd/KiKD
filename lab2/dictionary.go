package lab2

import (
	"math"
)

const symbolCount = 256
const precision uint8 = 32      //need to be able to represent at least 2^42
const whole uint64 = 4294967296 //2 ^ precision
const half uint64 = whole / 2
const quarter uint64 = half / 2

type dictionary struct {
	symbols    []symbol
	totalCount uint64
	minLength  uint8
	F          []uint64
}

func (d *dictionary) initialise() {
	*d = dictionary{
		symbols:    make([]symbol, symbolCount),
		totalCount: symbolCount,
		minLength:  11,
		F:          make([]uint64, symbolCount+1)}

	for i := 0; i < symbolCount; i++ {
		d.symbols[i] = symbol{
			code:  byte(i),
			count: 1}
	}

	for i := 0; i < symbolCount; i++ {
		f := uint64(0)

		for j := 0; j < i; j++ {
			f += d.symbols[j].count
		}

		d.F[i] = f
	}

	d.F[symbolCount] = d.totalCount
}

func (d *dictionary) rescale() {
	d.totalCount = 0

	for i := 0; i < symbolCount; i++ {
		d.symbols[i].count = (d.symbols[i].count + 1) / 2
		d.totalCount += d.symbols[i].count
	}

	for i := 0; i < symbolCount; i++ {
		f := uint64(0)

		for j := 0; j < i; j++ {
			f += d.symbols[j].count
		}

		d.F[i] = f
	}
}

func (d *dictionary) update(code byte) {
	i := int(code)
	minProb := 1.0

	d.symbols[i].count++
	d.totalCount++

	if d.totalCount >= 4*symbolCount {
		d.rescale()
	} else {
		for k := 0; k < symbolCount; k++ {
			f := uint64(0)

			for j := 0; j < k; j++ {
				f += d.symbols[j].count
			}

			d.F[k] = f
		}
	}

	for k := 0; k < symbolCount; k++ {
		if (float64(d.symbols[k].count) / float64(d.totalCount)) < minProb {
			minProb = (float64(d.symbols[k].count) / float64(d.totalCount))
		}
	}

	d.minLength = uint8(math.Log2(1/minProb)+1) + 3
}
