package lab2

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
)

var encoderDict dictionary
var encoderLeft uint64
var encoderRight uint64
var encoderCounter int
var encoderOutput string

func encode(b byte) {
	i := int(b)

	d := encoderRight - encoderLeft + 1
	encoderRight = encoderLeft + (encoderDict.F[i+1] * d / encoderDict.totalCount) - 1
	encoderLeft = encoderLeft + (encoderDict.F[i] * d / encoderDict.totalCount)
	encoderDict.update(b)

	for {
		if encoderRight < half || encoderLeft >= half {
			if encoderRight < half {
				encoderLeft *= 2
				encoderRight *= 2

				encoderOutput = encoderOutput + "0"
				for j := 0; j < encoderCounter; j++ {
					encoderOutput = encoderOutput + "1"
				}

				encoderCounter = 0
			} else {
				encoderLeft = 2 * (encoderLeft - half)
				encoderRight = 2 * (encoderRight - half)

				encoderOutput = encoderOutput + "1"
				for j := 0; j < encoderCounter; j++ {
					encoderOutput = encoderOutput + "0"
				}

				encoderCounter = 0
			}

			if encoderRight%2 == 0 {
				encoderRight++
			}
		} else if encoderLeft < half && encoderRight >= half && encoderLeft >= quarter && encoderRight <= 3*quarter {
			encoderLeft = 2 * (encoderLeft - quarter)
			encoderRight = 2 * (encoderRight - quarter)

			if encoderRight%2 == 0 {
				encoderRight++
			}

			encoderCounter++
		} else {
			break
		}
	}
}

func EncodeFile(in string, out string) {
	fileIn, errIn := os.Open(in)
	Check(errIn)
	defer fileIn.Close()

	fileOut, errOut := os.Create(out)
	Check(errOut)
	defer fileOut.Close()

	br := bufio.NewReader(fileIn)
	encoderDict.initialise()
	encoderLeft = 0
	encoderRight = whole
	encoderCounter = 0
	encoderOutput = ""

	inputSize := uint64(0)
	outputSize := uint64(0)

	for {
		b, e := br.ReadByte()

		if e != nil {
			if !errors.Is(e, io.EOF) {
				panic(e)
			} else {
				break
			}
		}

		inputSize++
		encode(b)

		for {
			if len(encoderOutput) >= 8 {
				outputByte := []byte{MakeByte(encoderOutput)}
				_, writeErr := fileOut.Write(outputByte)
				Check(writeErr)
				outputSize++
				encoderOutput = encoderOutput[8:]
			} else {
				break
			}
		}
	}

	tagString := TagToBitstring(encoderLeft, int(precision))

	encoderOutput = encoderOutput + string(tagString[0])

	for i := 0; i < encoderCounter; i++ {
		if tagString[0] == '0' {
			encoderOutput = encoderOutput + "1"
		} else {
			encoderOutput = encoderOutput + "0"
		}
	}

	for i := 1; i < len(tagString); i++ {
		encoderOutput = encoderOutput + string(tagString[i])
	}

	for {
		if len(encoderOutput) > 0 {
			outputByte := []byte{MakeByte(encoderOutput)}
			_, writeErr := fileOut.Write(outputByte)
			Check(writeErr)
			outputSize++
			encoderOutput = encoderOutput[int(math.Min(8.0, float64(len(encoderOutput)))):]
		} else {
			break
		}
	}

	fmt.Println("Entropy:", CalculateEntropy(&encoderDict))
	fmt.Println("Compression rate:", CalculateCompressionRate(inputSize, outputSize))
	fmt.Println("Average code length:", CalculateAverageCodeLength(inputSize, outputSize))
}
