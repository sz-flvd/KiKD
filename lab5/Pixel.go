package lab5

import (
	"math"
)

type pixel struct {
	r byte
	g byte
	b byte
}

func TaxicabDistance(x pixel, y pixel) int {
	dist := 0

	dist += int(math.Abs(float64(x.r) - float64(y.r)))
	dist += int(math.Abs(float64(x.g) - float64(y.g)))
	dist += int(math.Abs(float64(x.b) - float64(y.b)))

	return dist
}

func AveragePixel(pixels []pixel) pixel {
	var avg pixel
	r := 0.0
	g := 0.0
	b := 0.0
	l := float64(len(pixels))

	for _, p := range pixels {
		r += float64(p.r) / l
		g += float64(p.g) / l
		b += float64(p.b) / l
	}

	avg.r = byte(r)
	avg.g = byte(g)
	avg.b = byte(b)

	return avg
}

func (p *pixel) AddPerturbationVector() pixel {
	var q pixel

	q.r = (*p).r + 1
	q.g = (*p).g + 1
	q.b = (*p).b + 1

	return q
}
