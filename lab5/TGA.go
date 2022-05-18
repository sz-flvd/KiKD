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
	header    []byte
	imageData [][]pixel
	footer    []byte
	width     int
	height    int
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

	(*img).header = make([]byte, headerSize)
	(*img).footer = make([]byte, 0)

	for i = 0; i < headerSize; i++ {
		b[0], _ = br.ReadByte()
		(*img).header[i] = b[0]
	}

	(*img).width = int((*img).header[12]) + int((*img).header[13])*256
	(*img).height = int((*img).header[14]) + int((*img).header[15])*256

	(*img).imageData = make([][]pixel, (*img).height)

	for i = range (*img).imageData {
		(*img).imageData[i] = make([]pixel, (*img).width)
	}

	for i = range (*img).imageData {
		for j = range (*img).imageData[i] {
			b[0], _ = br.ReadByte()
			b[1], _ = br.ReadByte()
			b[2], _ = br.ReadByte()

			(*img).imageData[i][j] = pixel{
				r: b[0],
				g: b[1],
				b: b[2]}
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

		(*img).footer = append((*img).footer, b[0])
	}
}

func (img *TGAImage) CopyImage() TGAImage {
	copyImg := TGAImage{
		header:    make([]byte, len((*img).header)),
		imageData: make([][]pixel, len((*img).imageData)),
		footer:    make([]byte, len((*img).footer)),
		width:     (*img).width,
		height:    (*img).height}

	for i, col := range (*img).imageData {
		for range (*img).imageData {
			copyImg.imageData[i] = make([]pixel, len((*img).imageData[i]))
			copy(copyImg.imageData[i], col)
		}
	}

	copy(copyImg.header, (*img).header)
	copy(copyImg.footer, (*img).footer)

	return copyImg
}

func (img *TGAImage) SaveImage(filename string) {
	out, err := os.Create(filename)
	lab2.Check(err)
	defer out.Close()

	_, err = out.Write((*img).header)
	lab2.Check(err)

	for i := range (*img).imageData {
		for j := range (*img).imageData[i] {
			_, err = out.Write([]byte{(*img).imageData[i][j].r, (*img).imageData[i][j].g, (*img).imageData[i][j].b})
			lab2.Check(err)
		}
	}

	_, err = out.Write((*img).footer)
	lab2.Check(err)
}

func (img *TGAImage) FlattenImageData() []pixel {
	pixels := make([]pixel, 0)

	for _, p := range (*img).imageData {
		pixels = append(pixels, p...)
	}

	return pixels
}
