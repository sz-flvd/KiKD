package lab7

import (
	"KiKD/lab2"
	"bufio"
	"errors"
	"io"
	"os"
)

func Decode(in string, out string) {
	fileIn, errIn := os.Open(in)
	lab2.Check(errIn)
	defer fileIn.Close()

	fileOut, errOut := os.Create(out)
	lab2.Check(errOut)
	defer fileOut.Close()

	br := bufio.NewReader(fileIn)

	for {
		b, e := br.ReadByte()

		if e != nil {
			if !errors.Is(e, io.EOF) {
				lab2.Check(e)
			} else {
				break
			}
		}

		// do some decoding
	}

}
