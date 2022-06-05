package lab5

import (
	"KiKD/lab2"
	"bufio"
	"errors"
	"io"
	"os"
)

const (
	headerSize = 18
)

type TGAImage struct {
	Header    []byte
	ImageData [][]Pixel
	Footer    []byte
	Width     int
	Height    int
}

func (img *TGAImage) LoadImage(filename string) {
	var i int
	var j int
	b := make([]byte, 3)
	var e error

	in, err := os.Open(filename)
	lab2.Check(err)
	defer in.Close()

	br := bufio.NewReader(in)

	(*img).Header = make([]byte, headerSize)
	(*img).Footer = make([]byte, 0)

	for i = 0; i < headerSize; i++ {
		b[0], _ = br.ReadByte()
		(*img).Header[i] = b[0]
	}

	(*img).Width = int((*img).Header[12]) + int((*img).Header[13])*256
	(*img).Height = int((*img).Header[14]) + int((*img).Header[15])*256

	(*img).ImageData = make([][]Pixel, (*img).Height)

	for i = range (*img).ImageData {
		(*img).ImageData[i] = make([]Pixel, (*img).Width)
	}

	for i = range (*img).ImageData {
		for j = range (*img).ImageData[i] {
			b[0], _ = br.ReadByte()
			b[1], _ = br.ReadByte()
			b[2], _ = br.ReadByte()

			(*img).ImageData[i][j] = Pixel{
				R: b[0],
				G: b[1],
				B: b[2]}
		}
	}

	for {
		b[0], e = br.ReadByte()

		if e != nil {
			if !errors.Is(e, io.EOF) {
				lab2.Check(e)
			} else {
				break
			}
		}

		(*img).Footer = append((*img).Footer, b[0])
	}
}

func (img *TGAImage) CopyImage() TGAImage {
	copyImg := TGAImage{
		Header:    make([]byte, len((*img).Header)),
		ImageData: make([][]Pixel, len((*img).ImageData)),
		Footer:    make([]byte, len((*img).Footer)),
		Width:     (*img).Width,
		Height:    (*img).Height}

	for i, col := range (*img).ImageData {
		for range (*img).ImageData {
			copyImg.ImageData[i] = make([]Pixel, len((*img).ImageData[i]))
			copy(copyImg.ImageData[i], col)
		}
	}

	copy(copyImg.Header, (*img).Header)
	copy(copyImg.Footer, (*img).Footer)

	return copyImg
}

func (img *TGAImage) SaveImage(filename string) {
	out, err := os.Create(filename)
	lab2.Check(err)
	defer out.Close()

	_, err = out.Write((*img).Header)
	lab2.Check(err)

	for i := range (*img).ImageData {
		for j := range (*img).ImageData[i] {
			_, err = out.Write([]byte{(*img).ImageData[i][j].R, (*img).ImageData[i][j].G, (*img).ImageData[i][j].B})
			lab2.Check(err)
		}
	}

	_, err = out.Write((*img).Footer)
	lab2.Check(err)
}

func (img *TGAImage) FlattenImageData() []Pixel {
	pixels := make([]Pixel, 0)

	for _, p := range (*img).ImageData {
		pixels = append(pixels, p...)
	}

	return pixels
}
