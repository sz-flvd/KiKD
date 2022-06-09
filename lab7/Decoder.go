package lab7

import (
	"KiKD/lab2"
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
)

const parityErr = 32
const noErr = 64
const evenErr = 127

func Decode(in string, out string) {
	fmt.Printf("\t# Decoding [%s] to [%s]\n", in, out)

	fileIn, errIn := os.Open(in)
	lab2.Check(errIn)
	defer fileIn.Close()

	fileOut, errOut := os.Create(out)
	lab2.Check(errOut)
	defer fileOut.Close()

	br := bufio.NewReader(fileIn)
	errCount := 0
	tmpOutput := make([]byte, 0)
	decoderOutput := make([]byte, 0)

	for {
		b, e := br.ReadByte()

		if e != nil {
			if !errors.Is(e, io.EOF) {
				lab2.Check(e)
			} else {
				break
			}
		}

		if p := checkParity(b); p { // even number of errors occured [including no errors at all]
			if corrected := checkAndCorrect(b, p); corrected == noErr {
				tmpOutput = append(tmpOutput, byte(codeIndex(b)))
			} else if corrected == evenErr {
				errCount++
				tmpOutput = append(tmpOutput, byte(0))
			}
		} else { // odd number of errors occured and one of them can be corrected
			if correctedIndex := codeIndex(checkAndCorrect(b, p)); correctedIndex != -1 {
				tmpOutput = append(tmpOutput, byte(correctedIndex))
			} else {
				tmpOutput = append(tmpOutput, byte(0))
			}

		}
	}

	for i := 0; i < len(tmpOutput)-1; i += 2 {
		decoderOutput = append(decoderOutput, (tmpOutput[i]<<4)+tmpOutput[i+1])
	}

	numOutBytes, errWrite := fileOut.Write(decoderOutput)
	lab2.Check(errWrite)
	fmt.Printf("\t# %d bytes of decoded data written to [%s]\n", numOutBytes, out)
	fmt.Printf("\t# %d 2-bit errors found\n", errCount)

}

func checkParity(b byte) bool {
	parityBit := b % 2
	count := 0
	b >>= 1

	for {
		if b == 0 {
			break
		}

		count += int(b % 2)
		b >>= 1
	}

	return parityBit == byte(count%2)
}

func checkAndCorrect(b byte, parity bool) byte {
	var corrected byte
	var p1 byte
	var p2 byte
	var p3 byte
	b >>= 1
	bits := make([]byte, 0)

	for i := 0; i < 7; i++ {
		bits = append(bits, b%2)
		b >>= 1
	}

	p1 = (bits[0] + bits[1] + bits[2] + bits[4]) % 2
	p2 = (bits[1] + bits[2] + bits[3] + bits[5]) % 2
	p3 = (bits[2] + bits[3] + bits[4] + bits[6]) % 2

	errCode := (p1 << 2) + (p2 << 1) + p3

	if parity && errCode == 0 {
		corrected = noErr
	} else if parity && errCode != 0 {
		corrected = evenErr
	} else if !parity && errCode != 0 {
		errBit := (b >> (7 - errCode)) % 2
		if errBit == 0 {
			corrected = b + byte(math.Pow(2.0, float64(7-errCode)))
		} else { // errBit = 1
			corrected = b - byte(math.Pow(2.0, float64(7-errCode)))
		}
	} else {
		corrected = parityErr
	}

	return corrected
}

func codeIndex(b byte) int {
	for i, c := range Codes {
		if b == c {
			return i
		}
	}

	return -1
}
