package lab4

import (
	"KiKD/lab2"
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
	imageWidth = int(header[13])*256 + int(header[12])
	imageHeight = int(header[15])*256 + int(header[14])
	pixelDepth = int(header[16])

	bytes = make([][]byte, 3*imageHeight)

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

func Entropy() {

}
