package lab7

import (
	"KiKD/lab2"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

var Codes = [...]byte{0, 27, 53, 46, 105, 114, 92, 71, 209, 202, 228, 254, 184, 163, 141, 150}

func Encode(in string, out string) {
	fmt.Printf("\t# Encoding [%s] to [%s]\n", in, out)

	fileIn, errIn := os.Open(in)
	lab2.Check(errIn)
	defer fileIn.Close()

	fileOut, errOut := os.Create(out)
	lab2.Check(errOut)
	defer fileOut.Close()

	br := bufio.NewReader(fileIn)
	hammingOutput := make([]byte, 0)

	for {
		b, e := br.ReadByte()

		if e != nil {
			if !errors.Is(e, io.EOF) {
				lab2.Check(e)
			} else {
				break
			}
		}

		hammingOutput = append(hammingOutput, Codes[b>>4])
		hammingOutput = append(hammingOutput, Codes[b%16])
	}

	numOutBytes, errWrite := fileOut.Write(hammingOutput)
	lab2.Check(errWrite)
	fmt.Printf("\t# %d bytes of Hamming code written to [%s]\n", numOutBytes, out)
}
