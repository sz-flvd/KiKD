package lab2

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadAndEncode(filename string) {
	file, err := os.Open(filename)
	defer file.Close()
	check(err)

	br := bufio.NewReader(file)
	encoderDict.Initialise()

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
}

func ReadAndDecode(filename string) {
	decoderDict.Initialise()
}
