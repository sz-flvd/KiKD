package lab2

import (
	"math"
)

const symbolCount = 256

type dictionary struct {
	symbols    []symbol
	totalCount uint16
	minLength  uint8
	F          []float64
}

func (d *dictionary) initialise() {
	*d = dictionary{
		symbols:    make([]symbol, symbolCount),
		totalCount: symbolCount,
		minLength:  11,
		F:          make([]float64, symbolCount+1)}

	for i := 0; i < symbolCount; i++ {
		d.symbols[i] = symbol{
			code:        byte(i),
			count:       1,
			probability: 1.0 / symbolCount}
	}

	for i := 0; i < symbolCount; i++ {
		f := 0.0

		for j := 0; j < i; j++ {
			f += d.symbols[j].probability
		}

		d.F[i] = f
	}

	d.F[symbolCount] = 1.0
}

func (d *dictionary) rescale() {
	d.totalCount = 0

	for i := 0; i < symbolCount; i++ {
		d.symbols[i].count = (d.symbols[i].count + 1) / 2
		d.totalCount += d.symbols[i].count
	}

	for i := 0; i < symbolCount; i++ {
		d.symbols[i].probability = (float64)(d.symbols[i].count) / (float64)(d.totalCount)
	}

	for i := 0; i < symbolCount; i++ {
		f := 0.0

		for j := 0; j < i; j++ {
			f += d.symbols[j].probability
		}

		d.F[i] = f
	}
}

func (d *dictionary) update(code byte) {
	i := int(code)
	minProb := 1.0

	d.symbols[i].count++
	d.totalCount++

	if d.totalCount >= 2*symbolCount {
		d.rescale()
	} else {
		for k := 0; k < symbolCount; k++ {
			d.symbols[k].probability = (float64)(d.symbols[k].count) / (float64)(d.totalCount)
		}

		for k := 0; k < symbolCount; k++ {
			f := 0.0

			for j := 0; j < k; j++ {
				f += d.symbols[j].probability
			}

			d.F[k] = f
		}
	}

	for k := 0; k < symbolCount; k++ {
		if d.symbols[k].probability < minProb {
			minProb = d.symbols[k].probability
		}
	}

	d.minLength = uint8(math.Log2(1/minProb)+1) + 3
}
