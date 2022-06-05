package lab5

import (
	"math"
)

type Pixel struct {
	R byte
	G byte
	B byte
}

func TaxicabDistance(x Pixel, y Pixel) int {
	dist := 0

	dist += int(math.Abs(float64(x.R) - float64(y.R)))
	dist += int(math.Abs(float64(x.G) - float64(y.G)))
	dist += int(math.Abs(float64(x.B) - float64(y.B)))

	return dist
}

func AveragePixel(pixels []Pixel) Pixel {
	var avg Pixel
	r := 0.0
	g := 0.0
	b := 0.0
	l := float64(len(pixels))

	for _, p := range pixels {
		r += float64(p.R) / l
		g += float64(p.G) / l
		b += float64(p.B) / l
	}

	avg.R = byte(r)
	avg.G = byte(g)
	avg.B = byte(b)

	return avg
}

func (p *Pixel) AddPerturbationVector() Pixel {
	var q Pixel

	q.R = (*p).R + 1
	q.G = (*p).G + 1
	q.B = (*p).B + 1

	return q
}
