package lab7

import (
	"KiKD/lab2"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func Check(in1 string, in2 string) {
	fmt.Printf("\t# Comparing [%s] with [%s]\n", in1, in2)

	diffCount := 0

	fileIn1, errIn1 := os.Open(in1)
	lab2.Check(errIn1)
	defer fileIn1.Close()

	fileIn2, errIn2 := os.Open(in2)
	lab2.Check(errIn2)
	defer fileIn2.Close()

	br1 := bufio.NewReader(fileIn1)
	br2 := bufio.NewReader(fileIn2)

	for {
		b1, e1 := br1.ReadByte()
		b2, e2 := br2.ReadByte()

		if e1 != nil {
			if !errors.Is(e1, io.EOF) {
				lab2.Check(e1)
			} else {
				break
			}
		} else if e2 != nil {
			if !errors.Is(e2, io.EOF) {
				lab2.Check(e2)
			} else {
				break
			}
		}

		diffCount += compareBytes(b1, b2)
	}

	if diffCount == 0 {
		fmt.Printf("\t# Files [%s] and [%s] are identical!\n", in1, in2)
	} else {
		fmt.Printf("\t# %d different 4-bit blocks found between [%s] and [%s]\n", diffCount, in1, in2)
	}
}

func compareBytes(b1 byte, b2 byte) int {
	diffCount := 0

	for i := 0; i < 2; i++ {
		if b1%(1<<4) != b2%(1<<4) {
			diffCount++
		}

		b1 >>= 4
		b2 >>= 4
	}

	return diffCount
}
