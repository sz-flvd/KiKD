package lab4

import (
	"KiKD/lab2"
	"KiKD/lab3"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

const headerSize = 18

func LoadTGAFile(filename string) [][]byte {
	var imageWidth int
	var imageHeight int
	var IDLength int
	var colorMapType int
	var imageType int
	var pixelDepth int
	var i int
	var j int
	var b byte
	var e error
	var bytes [][]byte
	header := make([]byte, headerSize)
	in, err := os.Open(filename)

	lab2.Check(err)
	defer in.Close()

	br := bufio.NewReader(in)

	for i = 0; i < headerSize; i++ {
		b, _ = br.ReadByte()
		header[i] = b
	}

	IDLength = int(header[0])
	colorMapType = int(header[1])
	imageType = int(header[2])
	imageWidth = int(header[12]) + int(header[13])*256
	imageHeight = int(header[14]) + int(header[15])*256
	pixelDepth = int(header[16])

	bytes = make([][]byte, imageHeight)

	for i = range bytes {
		bytes[i] = make([]byte, 3*imageWidth)
	}

	for i = range bytes {
		for j = range bytes[i] {
			b, e = br.ReadByte()

			if e != nil {
				if !errors.Is(e, io.EOF) {
					lab2.Check(e)
				} else {
					break
				}
			}

			bytes[i][j] = b
		}
	}

	fmt.Println("----------", filename, " file information ----------")
	fmt.Println("ID length:", IDLength)
	fmt.Println("Color map type:", colorMapType)
	fmt.Println("Image type:", imageType)
	fmt.Println("Pixel depth:", pixelDepth)
	fmt.Println("Image width:", imageWidth)
	fmt.Println("Image height:", imageHeight)

	return bytes
}

func Entropy(bytes *[][]byte) (float64, float64, float64, float64) {
	total := make(map[byte]uint, 256)
	red := make(map[byte]uint, 256)
	green := make(map[byte]uint, 256)
	blue := make(map[byte]uint, 256)
	var entropy float64
	var entropyRed float64
	var entropyGreen float64
	var entropyBlue float64

	for i := range *bytes {
		for j := range (*bytes)[i] {
			total[(*bytes)[i][j]]++

			switch j % 3 {
			case 0:
				red[(*bytes)[i][j]]++
			case 1:
				green[(*bytes)[i][j]]++
			case 2:
				blue[(*bytes)[i][j]]++
			}
		}
	}

	entropy = lab3.Entropy(total)
	entropyRed = lab3.Entropy(red)
	entropyGreen = lab3.Entropy(green)
	entropyBlue = lab3.Entropy(blue)

	return entropy, entropyRed, entropyGreen, entropyBlue
}

func IntArrayToByteArray(arr *[][]int) [][]byte {
	bytes := make([][]byte, len(*arr))

	for i := range bytes {
		bytes[i] = make([]byte, len((*arr)[i]))
	}

	for i := range *arr {
		for j := range (*arr)[i] {
			bytes[i][j] = byte((*arr)[i][j])
		}
	}

	return bytes
}
