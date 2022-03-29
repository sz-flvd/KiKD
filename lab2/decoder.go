package lab2

import (
	"bufio"
	"errors"
	"io"
	"math"
	"os"
)

var decoderDict dictionary
var decoderLeft float64
var decoderRight float64
var decoderCounter int
var decoderInput string
var decoderOutput []byte

func decode() {
	for {
		tag := getTagValue(decoderInput[:int(math.Min(float64(len(decoderInput)), float64(decoderDict.minLength)))])
		d := decoderRight - decoderLeft
		for i := 0; i < symbolCount; i++ {
			if i < 255 {
				if decoderLeft+decoderDict.symbols[i].F*d <= tag && tag < decoderLeft+decoderDict.symbols[i+1].F*d {
					decoderOutput = append(decoderOutput, decoderDict.symbols[i].code)
					decoderRight = decoderLeft + decoderDict.symbols[i+1].F*d
					decoderLeft = decoderLeft + decoderDict.symbols[i].F*d
					decoderDict.update(byte(i))
					break
				}
			} else {
				decoderOutput = append(decoderOutput, decoderDict.symbols[len(decoderDict.symbols)-1].code)
				decoderRight = decoderLeft + d
				decoderLeft = decoderLeft + decoderDict.symbols[i].F*d
				decoderDict.update(byte(i))
			}
		}

		for {
			if decoderLeft >= 0 && decoderRight <= 0.5 {
				decoderLeft *= 2
				decoderRight *= 2

				for j := 0; j < decoderCounter+1; j++ {
					if len(decoderInput) == 0 {
						break
					}

					decoderInput = decoderInput[1:]
				}
			} else if decoderLeft >= 0.5 && decoderRight <= 1 {
				decoderLeft = 2*decoderLeft - 1
				decoderRight = 2*decoderRight - 1

				for j := 0; j < decoderCounter+1; j++ {
					if len(decoderInput) == 0 {
						break
					}

					decoderInput = decoderInput[1:]
				}
			} else if decoderLeft < 0.5 && decoderRight > 0.5 && decoderLeft >= 0.25 && decoderRight <= 0.75 {
				decoderLeft = 2*decoderLeft - 0.5
				decoderRight = 2*decoderRight - 0.5
				decoderCounter++
			} else {
				break
			}
		}

		if len(decoderInput) <= int(decoderDict.minLength) {
			break
		}
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
		b, e := br.ReadByte()

		if e != nil && !errors.Is(e, io.EOF) {
			panic(e)
		}

		if e != nil {
			break
		}

		decoderInput = decoderInput + makeBitstring(b)
	}

	decode()

	_, writeErr := fileOut.Write(decoderOutput)
	check(writeErr)
	decoderOutput = decoderOutput[:0]
}
