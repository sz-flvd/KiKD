package lab7

import (
	"KiKD/lab2"
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func GenerateNoise(p float64, in string, out string) {
	fmt.Printf("\t# Generating noise from [%s] to [%s] (with %f probability for each bit)\n", in, out, p)

	fileIn, errIn := os.Open(in)
	lab2.Check(errIn)
	defer fileIn.Close()

	fileOut, errOut := os.Create(out)
	lab2.Check(errOut)
	defer fileOut.Close()

	br := bufio.NewReader(fileIn)
	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)

	for {
		b, e := br.ReadByte()

		if e != nil {
			if !errors.Is(e, io.EOF) {
				lab2.Check(e)
			} else {
				break
			}
		}

		_, errWrite := fileOut.Write([]byte{shiftBits(b, p, random)})
		lab2.Check(errWrite)
	}
}

func shiftBits(b byte, p float64, random *rand.Rand) byte {
	var out byte
	var mod byte

	for i := 0; i < 8; i++ {
		mod = b % 2
		if r := random.Float64(); r < p {
			mod = 1 - mod
		}
		mod <<= i
		out += mod
		b >>= 1
	}

	return out
}
