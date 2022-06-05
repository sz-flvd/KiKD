package lab7

import (
	"KiKD/lab2"
	"bufio"
	"errors"
	"io"
	"math/rand"
	"os"
	"time"
)

func GenerateNoise(p float64, in string, out string) {
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

		_, errWrite := fileOut.Write([]byte{shiftBits(b, p)})
		lab2.Check(errWrite)
	}
}

func shiftBits(b byte, p float64) byte {
	var out byte
	var mod byte

	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)

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
