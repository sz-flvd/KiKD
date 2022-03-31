package lab2

import (
	"bufio"
	"errors"
	"io"
	"os"
)

var decoderDict dictionary
var decoderLeft uint64
var decoderRight uint64
var decoderInput string
var decoderOutput []byte

func decode() {
	tag := GetTagValue(decoderInput[:precision], precision)
	decoderInput = decoderInput[precision:]

	for {
		for i := 0; i < symbolCount; i++ {
			d := decoderRight - decoderLeft
			if decoderLeft+(decoderDict.F[i]*d/decoderDict.totalCount) <= tag && tag < decoderLeft+(decoderDict.F[i+1]*d/decoderDict.totalCount) {
				decoderOutput = append(decoderOutput, decoderDict.symbols[i].code)
				decoderRight = decoderLeft + (decoderDict.F[i+1] * d / decoderDict.totalCount)
				decoderLeft = decoderLeft + (decoderDict.F[i] * d / decoderDict.totalCount)
				decoderDict.update(byte(i))
				break
			}
		}

		for {
			if decoderRight <= half || decoderLeft >= half {
				if decoderRight <= half {
					decoderLeft *= 2
					decoderRight = 2*decoderRight + 1
					tag *= 2
				} else {
					decoderLeft = 2 * (decoderLeft - half)
					decoderRight = 2*(decoderRight-half) + 1
					tag = 2 * (tag - half)
				}

				if decoderInput[0] == '1' {
					tag += 1
				}
				decoderInput = decoderInput[1:]
			} else if decoderLeft < half && decoderRight > half && decoderLeft >= quarter && decoderRight <= 3*quarter {
				decoderLeft = 2 * (decoderLeft - quarter)
				decoderRight = 2*(decoderRight-quarter) + 1
				tag = 2 * (tag - quarter)

				if decoderInput[0] == '1' {
					tag += 1
				}
				decoderInput = decoderInput[1:]
			} else {
				break
			}
		}

		if len(decoderInput) < int(precision) {
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
	decoderLeft = 0
	decoderRight = whole
	decoderInput = ""
	decoderOutput = make([]byte, 0)

	for {
		b, e := br.ReadByte()

		if e != nil {
			if !errors.Is(e, io.EOF) {
				Check(e)
			} else {
				break
			}
		}

		decoderInput = decoderInput + MakeBitstring(b)
	}

	for i := 0; i < int(precision); i++ {
		decoderInput = decoderInput + "0"
	}

	decode()

	decoderOutput = decoderOutput[:len(decoderOutput)-1]

	_, writeErr := fileOut.Write(decoderOutput)
	Check(writeErr)
}
