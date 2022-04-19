package lab3

import (
	"KiKD/lab2"
	"bufio"
	"errors"
	"io"
	"os"
)

func decode(code string, param int) (int, int) {
	switch param {
	case Gamma:
		return eliasGammaDecode(code)
	case Delta:
		return eliasDeltaDecode(code)
	case Omega:
		return eliasOmegaDecode(code)
	case Fib:
		return fibonacciDecode(code)
	default:
		return 0, 0
	}
}

func Decode(in string, out string, param int) {
	fileIn, errIn := os.Open(in)
	fileOut, errOut := os.Create(out)

	br := bufio.NewReader(fileIn)
	var dict dictionary
	var b byte
	var e error
	var code int
	var length int
	input := ""

	lab2.Check(errIn)
	defer fileIn.Close()

	lab2.Check(errOut)
	defer fileOut.Close()

	dict.initialise()

	for {
		b, e = br.ReadByte()

		if e != nil {
			if !errors.Is(e, io.EOF) {
				lab2.Check(e)
			} else {
				break
			}
		}

		input = input + lab2.MakeBitstring(b)
		code, length = decode(input, param)

		if code != 0 && length != 0 {
			_, writeErr := fileOut.Write(dict.entries[code-1].bytes)
			lab2.Check(writeErr)
			// outputSize++
			if length > len(input)-1 {
				input = input[:0]
			} else {
				input = input[length:]
			}

			if dict.last > symbolCount {
				dict.entries[len(dict.entries)-1].bytes = append(dict.entries[len(dict.entries)-1].bytes, dict.entries[code-1].bytes[0])
			}
			dict.newEntry(dict.entries[code-1].bytes)
		}
	}
}
