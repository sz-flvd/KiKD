package lab2

import (
	"bufio"
	"errors"
	"io"
	"os"
)

var decoderDict dictionary
var decoderLeft float64
var decoderRight float64
var decoderMid float64
var decoderCounter int
var decoderInput string
var decoderOutput []byte

func decode() {
	for {
		tag := GetTagValue(decoderInput[:decoderDict.minLength])

		for i := 0; i < symbolCount; i++ {
			if decoderLeft+decoderDict.F[i]*(decoderRight-decoderLeft) <= tag && tag < decoderLeft+decoderDict.F[i+1]*(decoderRight-decoderLeft) {
				decoderOutput = append(decoderOutput, decoderDict.symbols[i].code)
				decoderMid = decoderRight - decoderLeft
				decoderRight = decoderLeft + decoderDict.F[i+1]*decoderMid
				decoderLeft = decoderLeft + decoderDict.F[i]*decoderMid
				decoderDict.update(byte(i))
				break
			}
		}

		for {
			if decoderLeft >= 0.0 && decoderRight <= 0.5 {
				decoderLeft *= 2.0
				decoderRight *= 2.0

				if len(decoderInput) <= decoderCounter {
					decoderInput = decoderInput[:0]
					break
				}

				decoderInput = decoderInput[decoderCounter+1:]
				decoderCounter = 0
			} else if decoderLeft >= 0.5 && decoderRight <= 1.0 {
				decoderLeft = 2.0*decoderLeft - 1.0
				decoderRight = 2.0*decoderRight - 1.0

				if len(decoderInput) <= decoderCounter {
					decoderInput = decoderInput[:0]
					break
				}

				decoderInput = decoderInput[decoderCounter+1:]
				decoderCounter = 0
			} else if decoderLeft < 0.5 && decoderRight > 0.5 && decoderLeft >= 0.25 && decoderRight <= 0.75 {
				decoderLeft = 2.0*decoderLeft - 0.5
				decoderRight = 2.0*decoderRight - 0.5
				decoderCounter++
			} else {
				break
			}
		}

		if len(decoderInput) < int(decoderDict.minLength) {
			break
		}
	}
}

func DecodeFile(in string, out string) {
	fileIn, errIn := os.Open(in)
	Check(errIn)
	defer fileIn.Close()

	fileOut, errOut := os.Create(out)
	Check(errOut)
	defer fileOut.Close()

	br := bufio.NewReader(fileIn)
	decoderDict.initialise()
	decoderLeft = 0.0
	decoderRight = 1.0
	decoderMid = 0.5
	decoderCounter = 0
	decoderInput = ""
	decoderOutput = make([]byte, 0)

	for {
		b, e := br.ReadByte()

		if e != nil {
			if !errors.Is(e, io.EOF) {
				panic(e)
			} else {
				break
			}
		}

		decoderInput = decoderInput + MakeBitstring(b)
	}

	for {
		if decoderInput[len(decoderInput)-1] == '0' {
			decoderInput = decoderInput[:len(decoderInput)-1]
		} else {
			break
		}
	}

	decode()

	_, writeErr := fileOut.Write(decoderOutput)
	Check(writeErr)
	decoderOutput = decoderOutput[:0]
}
