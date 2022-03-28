package lab2

import (
	"bufio"
	"errors"
	"fmt"
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
	if i == 255 {
		encoderRight = encoderLeft + d
	} else {
		encoderRight = encoderLeft + encoderDict.symbols[i+1].F*d
	}
	encoderLeft = encoderLeft + encoderDict.symbols[i].F*d

	for {
		if encoderLeft >= 0 && encoderRight <= 0.5 {
			encoderLeft *= 2
			encoderRight *= 2

			encoderOutput = encoderOutput + "0"
			for j := 0; j < encoderCounter; j++ {
				encoderOutput = encoderOutput + "1"
			}

			encoderCounter = 0
		} else if encoderLeft >= 0.5 && encoderRight <= 1 {
			encoderLeft = 2*encoderLeft - 1
			encoderRight = 2*encoderRight - 1

			encoderOutput = encoderOutput + "1"
			for j := 0; j < encoderCounter; j++ {
				encoderOutput = encoderOutput + "0"
			}

			encoderCounter = 0
		} else if encoderLeft < 0.5 && encoderRight > 0.5 && encoderLeft >= 0.25 && encoderRight <= 0.75 {
			encoderLeft = 2*encoderLeft - 0.5
			encoderRight = 2*encoderRight - 0.5
			encoderCounter++
		} else {
			break
		}
	}

	encoderDict.update(b)

	if encoderDict.totalCount >= 2*symbolCount {
		encoderDict.rescale()
	}
}

func EncodeFile(in string, out string) {
	fileIn, errIn := os.Open(in)
	check(errIn)
	defer fileIn.Close()

	fileOut, errOut := os.Create(out)
	check(errOut)
	defer fileOut.Close()

	br := bufio.NewReader(fileIn)
	encoderDict.initialise()
	encoderLeft = 0
	encoderRight = 1
	encoderCounter = 0
	encoderOutput = ""

	for {
		b, e := br.ReadByte()

		if e != nil && !errors.Is(e, io.EOF) {
			fmt.Println(e)
			break
		}

		encode(b)

		for {
			if len(encoderOutput) >= 8 {
				byte := []byte{makeByte(encoderOutput)}
				_, writeErr := fileOut.Write(byte)
				check(writeErr)
				encoderOutput = encoderOutput[8:]
			} else {
				break
			}
		}

		if e != nil {
			break
		}
	}

	for {
		if len(encoderOutput) >= 8 {
			byte := []byte{makeByte(encoderOutput)}
			_, writeErr := fileOut.Write(byte)
			check(writeErr)
			encoderOutput = encoderOutput[8:]
		} else {
			break
		}
	}

}
