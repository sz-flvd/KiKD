package lab2

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

var decoderDict dictionary
var decoderLeft float64
var decoderRight float64
var decoderCounter int
var decoderInput string
var decoderOutput []byte

func decode(bitstring string) {
	if decoderLeft < decoderRight {
		fmt.Println(decoderCounter) // przestań krzyczeć na mnie że nie używam tych zmiennych pls
	}
}

func DecodeFile(in string, out string) {
	fileIn, errIn := os.Open(in)
	check(errIn)
	defer fileIn.Close()

	fileOut, errOut := os.Create(out)
	check(errOut)
	defer fileOut.Close()

	br := bufio.NewReader(fileIn)
	decoderDict.initialise()
	decoderLeft = 0.0
	decoderRight = 1.0
	decoderCounter = 0
	decoderInput = ""
	decoderOutput = make([]byte, 0)

	for {
		var e error
		for {
			b, e := br.ReadByte()

			if e != nil && !errors.Is(e, io.EOF) {
				panic(e)
			}

			decoderInput = decoderInput + makeBitstring(b)

			if len(decoderInput) >= int(decoderDict.minLength) || e != nil {
				break
			}
		}

		decode(decoderInput)

		_, writeErr := fileOut.Write(decoderOutput)
		check(writeErr)
		decoderOutput = decoderOutput[:0]

		if e != nil {
			break
		}
	}
}
