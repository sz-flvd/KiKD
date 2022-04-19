package lab3

import (
	"KiKD/lab2"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	Gamma = iota
	Delta = iota
	Omega = iota
	Fib   = iota
)

func encode(num int, param int) string {
	switch param {
	case Gamma:
		return eliasGamma(num)
	case Delta:
		return eliasDelta(num)
	case Omega:
		return eliasOmega(num)
	case Fib:
		return fibonacci(num)
	default:
		return ""
	}
}

func Encode(in string, out string, param int) {
	fileIn, errIn := os.Open(in)
	lab2.Check(errIn)
	defer fileIn.Close()
	fileInStat, errInStat := fileIn.Stat()
	lab2.Check(errInStat)

	fileOut, errOut := os.Create(out)
	lab2.Check(errOut)
	defer fileOut.Close()

	br := bufio.NewReader(fileIn)
	var dict dictionary
	currBytes := make([]byte, 0)
	symbols := make(map[byte]uint)
	var b byte
	var e error
	var outputSize uint = 0
	output := ""

	dict.initialise()

	for {
		b, _ = br.ReadByte()

		symbols[b]++
		currBytes = append(currBytes, b)

		_, e = br.Peek(1)

		if e != nil {
			if !errors.Is(e, io.EOF) {
				lab2.Check(e)
			} else {
				break
			}
		}

		currBytes = append(currBytes, first(br.Peek(1))[0])

		if dict.find(currBytes) == -1 {
			dict.newEntry(currBytes)
			currBytes = currBytes[:len(currBytes)-1]
			output += encode(dict.find(currBytes)+1, param)

			for {
				if len(output) >= 8 {
					outputByte := []byte{lab2.MakeByte(output)}
					_, writeErr := fileOut.Write(outputByte)
					lab2.Check(writeErr)
					outputSize++
					output = output[8:]
				} else {
					break
				}
			}

			currBytes = currBytes[:0]
		} else {
			currBytes = currBytes[:len(currBytes)-1]
		}
	}

	output += encode(dict.find(currBytes)+1, param)

	for {
		if len(output) >= 8 {
			outputByte := []byte{lab2.MakeByte(output)}
			_, writeErr := fileOut.Write(outputByte)
			lab2.Check(writeErr)
			outputSize++
			output = output[8:]
		} else {
			break
		}
	}

	outputByte := []byte{lab2.MakeByte(output)}
	_, writeErr := fileOut.Write(outputByte)
	lab2.Check(writeErr)
	outputSize++

	fmt.Printf("Size of file before compression: %d\n", fileInStat.Size())
	fmt.Printf("Size of file after compression: %d\n", outputSize)
	fmt.Printf("Compression rate: %f\n", lab2.CalculateCompressionRate(uint64(fileInStat.Size()), uint64(outputSize)))
	fmt.Printf("Entropy of input file: %f\n", entropy(symbols))
	fmt.Printf("Entropy of output code: \n")
}
