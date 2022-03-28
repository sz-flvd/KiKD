package lab1

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
)

type pair struct {
	first  string
	second string
}

func probability(bitstrings []string) map[string]float64 {
	p := make(map[string]float64)

	for _, b := range bitstrings {

		if value, found := p[b]; found {
			p[b] = value + 1
		} else {
			p[b] = 1
		}
	}

	for k := range p {
		p[k] /= float64(len(bitstrings))
	}

	return p
}

func conditionalProbability(bitstrings []string, p map[string]float64) map[pair]float64 {
	cp := make(map[pair]float64)

	previous := "0"

	for _, b := range bitstrings {
		pair := pair{previous, b}
		if value, found := cp[pair]; found {
			cp[pair] = value + 1
		} else {
			cp[pair] = 1
		}
		previous = b
	}

	for k := range cp {
		cp[k] /= p[k.first] * float64(len(bitstrings))
	}

	return cp
}

func entropy(p map[string]float64) float64 {
	e := 0.0

	for _, v := range p {
		e += v * -math.Log2(v)
	}

	return e
}

func conditionalEntropy(p map[string]float64, cp map[pair]float64) float64 {
	ce := 0.0

	for k := range p {
		h := 0.0

		for i := range p {
			if v, f := cp[pair{k, i}]; f {
				h += v * -math.Log2(v)
			}
		}

		ce += p[k] * h
	}

	return ce
}

func readFile(filename string) []string {
	bytes, _ := ioutil.ReadFile(filename)

	bitstrings := make([]string, 0)

	for _, b := range bytes {
		bitstring := strconv.FormatUint(uint64(b), 2)
		l := len(bitstring)
		for i := 0; i < 8-l; i++ {
			bitstring = "0" + bitstring
		}
		bitstrings = append(bitstrings, bitstring)
	}

	return bitstrings
}

func ReadAndAnalyse(filename string) {
	bitstrings := readFile(filename)

	p := probability(bitstrings)
	fmt.Println(len(p))

	cp := conditionalProbability(bitstrings, p)
	fmt.Println(len(cp))

	e := entropy(p)
	fmt.Println("Entropy:", e)

	ce := conditionalEntropy(p, cp)
	fmt.Println("Conditional entropy:", ce)

	r := math.Abs(ce - e)
	fmt.Println("Difference:", r)
}
