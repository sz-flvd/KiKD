package lab3

import (
	"KiKD/lab2"
	"os"
)

const (
	Gamma = iota
	Delta = iota
	Omega = iota
	Fib   = iota
)

func Encode(in string, out string, param int) {
	fileIn, errIn := os.Open(in)
	lab2.Check(errIn)
	defer fileIn.Close()

	fileOut, errOut := os.Create(out)
	lab2.Check(errOut)
	defer fileOut.Close()
}
