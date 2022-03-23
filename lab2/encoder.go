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
var output []int

func encode(b byte) {
	i := int(b)

	d := encoderRight - encoderLeft
	if i == 255 {
		encoderRight = encoderLeft + d
	} else {
		encoderRight = encoderLeft + encoderDict.symbols[i+1].F*d
	}
	encoderLeft = encoderLeft + encoderDict.symbols[i].F*d

	if encoderLeft >= 0 && encoderRight <= 0.5 {
		encoderLeft *= 2
		encoderRight *= 2

		output = append(output, 0) //change to outputting to file
		for j := 0; j < encoderCounter; j++ {
			output = append(output, 1) //change to outputting to file
		}

		encoderCounter = 0
	} else if encoderLeft >= 0.5 && encoderRight <= 1 {
		encoderLeft = 2*encoderLeft - 1
		encoderRight = 2*encoderRight - 1

		output = append(output, 1) //change to outputting to file
		for j := 0; j < encoderCounter; j++ {
			output = append(output, 0) //change to outputting to file
		}

		encoderCounter = 0
	} else if encoderLeft < 0.5 && encoderRight > 0.5 && encoderLeft >= 0.25 && encoderRight <= 0.75 {
		encoderLeft = 2*encoderLeft - 0.5
		encoderRight = 2*encoderRight - 0.5
		encoderCounter++
	}

	encoderDict.update(b)

	if encoderDict.totalCount >= 2*symbolCount {
		encoderDict.rescale()
	}
}

func EncodeFile(filename string) {
	file, err := os.Open(filename)
	defer file.Close()
	check(err)

	br := bufio.NewReader(file)
	encoderDict.initialise()
	encoderLeft = 0
	encoderRight = 1
	encoderCounter = 0
	output = make([]int, 0)

	for {
		b, e := br.ReadByte()

		if e != nil && !errors.Is(e, io.EOF) {
			fmt.Println(e)
			break
		}

		encode(b)

		if e != nil {
			break
		}
	}

	fmt.Println(output) //close output file
}
