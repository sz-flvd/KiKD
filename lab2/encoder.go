package lab2

import (
	"bufio"
	"errors"
	"io"
	"os"
)

var encoderDict dictionary
var encoderLeft float64
var encoderRight float64
var encoderCounter int
var encoderOutput string

func encode(b byte) {
	i := int(b)

	d := encoderRight - encoderLeft
	encoderRight = encoderLeft + encoderDict.F[i+1]*d
	encoderLeft = encoderLeft + encoderDict.F[i]*d
	encoderDict.update(b)

	for {
		if encoderLeft >= 0.0 && encoderRight <= 0.5 {
			encoderLeft *= 2.0
			encoderRight *= 2.0

			encoderOutput = encoderOutput + "0"
			for j := 0; j < encoderCounter; j++ {
				encoderOutput = encoderOutput + "1"
			}

			encoderCounter = 0
		} else if encoderLeft >= 0.5 && encoderRight <= 1.0 {
			encoderLeft = 2.0*encoderLeft - 1.0
			encoderRight = 2.0*encoderRight - 1.0

			encoderOutput = encoderOutput + "1"
			for j := 0; j < encoderCounter; j++ {
				encoderOutput = encoderOutput + "0"
			}

			encoderCounter = 0
		} else if encoderLeft < 0.5 && encoderRight > 0.5 && encoderLeft >= 0.25 && encoderRight <= 0.75 {
			encoderLeft = 2.0*encoderLeft - 0.5
			encoderRight = 2.0*encoderRight - 0.5
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
	encoderLeft = 0.0
	encoderRight = 1.0
	encoderCounter = 0
	encoderOutput = ""

	for {
		b, e := br.ReadByte()

		if e != nil {
			if !errors.Is(e, io.EOF) {
				panic(e)
			} else {
				break
			}
		}

		encode(b)

		for {
			if len(encoderOutput) >= 8 {
				outputByte := []byte{MakeByte(encoderOutput)}
				_, writeErr := fileOut.Write(outputByte)
				Check(writeErr)
				encoderOutput = encoderOutput[8:]
			} else {
				break
			}
		}
	}

	encoderOutput = encoderOutput + FloatToBitstring((encoderLeft+encoderRight)/2.0, int(encoderDict.minLength))

	for {
		if len(encoderOutput) >= 8 {
			outputByte := []byte{MakeByte(encoderOutput)}
			_, writeErr := fileOut.Write(outputByte)
			Check(writeErr)
			encoderOutput = encoderOutput[8:]
		} else {
			break
		}
	}

	outputByte := []byte{MakeByte(encoderOutput)}
	_, writeErr := fileOut.Write(outputByte)
	Check(writeErr)
}
