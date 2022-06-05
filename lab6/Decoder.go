package lab6

import (
	"KiKD/lab2"
	"KiKD/lab5"
	"bufio"
	"errors"
	"io"
	"os"
)

func Decode(filenameIn string, filenameOut string) {
	var b byte
	var e error
	var footerLen int
	var quantizerBits int

	var img lab5.TGAImage

	in, err := os.Open(filenameIn)
	lab2.Check(err)
	defer in.Close()

	br := bufio.NewReader(in)

	b, _ = br.ReadByte()
	footerLen = int(b)

	b, _ = br.ReadByte()
	quantizerBits = int(b)

	quantizerRL := make([]int, 0)
	quantizerRH := make([]int, 0)
	quantizerGL := make([]int, 0)
	quantizerGH := make([]int, 0)
	quantizerBL := make([]int, 0)
	quantizerBH := make([]int, 0)

	b, _ = br.ReadByte()
	usedRL := int(b)
	for i := 0; i < usedRL; i++ {
		b, _ = br.ReadByte()
		quantizerRL = append(quantizerRL, int(b))
	}

	b, _ = br.ReadByte()
	usedRH := int(b)
	for i := 0; i < usedRH; i++ {
		b, _ = br.ReadByte()
		quantizerRH = append(quantizerRH, int(b)-128)
	}

	b, _ = br.ReadByte()
	usedGL := int(b)
	for i := 0; i < usedGL; i++ {
		b, _ = br.ReadByte()
		quantizerGL = append(quantizerGL, int(b))
	}

	b, _ = br.ReadByte()
	usedGH := int(b)
	for i := 0; i < usedGH; i++ {
		b, _ = br.ReadByte()
		quantizerGH = append(quantizerGH, int(b)-128)
	}

	b, _ = br.ReadByte()
	usedBL := int(b)
	for i := 0; i < usedBL; i++ {
		b, _ = br.ReadByte()
		quantizerBL = append(quantizerBL, int(b))
	}

	b, _ = br.ReadByte()
	usedBH := int(b)
	for i := 0; i < usedBH; i++ {
		b, _ = br.ReadByte()
		quantizerBH = append(quantizerBH, int(b)-128)
	}

	img.Header = make([]byte, 18)
	for i := range img.Header {
		b, _ := br.ReadByte()
		img.Header[i] = b
	}
	img.Width = int(img.Header[12]) + int(img.Header[13])*256
	img.Height = int(img.Header[14]) + int(img.Header[15])*256

	img.Footer = make([]byte, footerLen)
	for i := range img.Footer {
		b, _ := br.ReadByte()
		img.Footer[i] = b
	}

	tmpImageData := make([]byte, 0)

	for {
		b, e = br.ReadByte()

		if e != nil {
			if !errors.Is(e, io.EOF) {
				lab2.Check(e)
			} else {
				break
			}
		}

		tmpImageData = append(tmpImageData, b)
	}

	seqLen := len(tmpImageData) / 6
	bytesRL := tmpImageData[0:seqLen]
	bytesRH := tmpImageData[seqLen : 2*seqLen]
	bytesGL := tmpImageData[2*seqLen : 3*seqLen]
	bytesGH := tmpImageData[3*seqLen : 4*seqLen]
	bytesBL := tmpImageData[4*seqLen : 5*seqLen]
	bytesBH := tmpImageData[5*seqLen : 6*seqLen]

	seqRL := decompress(bytesRL, quantizerRL, quantizerBits)
	seqRH := decompress(bytesRH, quantizerRH, quantizerBits)
	seqGL := decompress(bytesGL, quantizerGL, quantizerBits)
	seqGH := decompress(bytesGH, quantizerGH, quantizerBits)
	seqBL := decompress(bytesBL, quantizerBL, quantizerBits)
	seqBH := decompress(bytesBH, quantizerBH, quantizerBits)

	red := make([]int, 0)
	green := make([]int, 0)
	blue := make([]int, 0)

	for i := range seqRL {
		red = append(red, seqRL[i])
		red = append(red, seqRH[i])
		green = append(green, seqGL[i])
		green = append(green, seqGH[i])
		blue = append(blue, seqBL[i])
		blue = append(blue, seqBH[i])
	}

	for i := 0; i < len(red)-1; i += 2 {
		red[i], red[i+1] = undoFilter(red[i], red[i+1])
		green[i], green[i+1] = undoFilter(green[i], green[i+1])
		blue[i], blue[i+1] = undoFilter(blue[i], blue[i+1])
	}

	img.ImageData = make([][]lab5.Pixel, img.Height)

	for i := range img.ImageData {
		img.ImageData[i] = make([]lab5.Pixel, img.Width)
	}

	for i := 0; i < img.Height; i++ {
		for j := 0; j < img.Width; j++ {
			img.ImageData[i][j].R = byte(red[i*img.Width+j])
			img.ImageData[i][j].G = byte(green[i*img.Width+j])
			img.ImageData[i][j].B = byte(blue[i*img.Width+j])
		}
	}

	// the end -> save file
	(&img).SaveImage(filenameOut)
}

func undoFilter(y int, z int) (int, int) {
	x2 := y + z
	x1 := 2*y - x2

	return x1, x2
}

func decompress(data []byte, quantDict []int, bitrate int) []int {
	decompressedData := make([]int, 0)
	stringData := ""

	for _, b := range data {
		stringData += byteToString(b, 8)
	}

	stringData = stringData[:len(stringData)-(len(stringData)%bitrate)]

	for i := 0; i < len(stringData)-bitrate+1; i += bitrate {
		decompressedData = append(decompressedData, int(stringToByte(stringData[i:i+bitrate])))
	}

	for i := range decompressedData {
		//fmt.Println(i, decompressedData[i], len(quantDict))
		decompressedData[i] = quantDict[decompressedData[i]]
	}

	return decompressedData
}
