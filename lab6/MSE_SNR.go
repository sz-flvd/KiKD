package lab6

import (
	"KiKD/lab5"
	"fmt"
	"math"
)

func ShowMSESNR(originalFilename string, compressedFilename string) {
	fmt.Printf("Showing MSE and SNR between %s and %s\n", originalFilename, compressedFilename)

	var originalImg lab5.TGAImage
	var compressedImg lab5.TGAImage

	(&originalImg).LoadImage(originalFilename)
	(&compressedImg).LoadImage(compressedFilename)

	mse, mseR, mseG, mseB := MSE(&originalImg, &compressedImg)
	snr, snrR, snrG, snrB := SNR(&originalImg, mse, mseR, mseG, mseB)

	fmt.Println("MSE")
	fmt.Printf("Total:\t%f\n", mse)
	fmt.Printf("Red:\t%f\n", mseR)
	fmt.Printf("Green:\t%f\n", mseG)
	fmt.Printf("Blue:\t%f\n", mseB)
	fmt.Println("SNR\t")
	fmt.Printf("Total:\t%f\t[%f dB]\n", snr, 10*math.Log10(snr))
	fmt.Printf("Red:\t%f\t[%f dB]\n", snrR, 10*math.Log10(snrR))
	fmt.Printf("Green:\t%f\t[%f dB]\n", snrG, 10*math.Log10(snrG))
	fmt.Printf("Blue:\t%f\t[%f dB]\n", snrB, 10*math.Log10(snrB))
}

func MSE(originalImg *lab5.TGAImage, compressedImg *lab5.TGAImage) (float64, float64, float64, float64) {
	var mse float64 = 0.0
	var mseR float64 = 0.0
	var mseG float64 = 0.0
	var mseB float64 = 0.0

	for i, row := range (*originalImg).ImageData {
		for j, p := range row {
			r := math.Pow(float64(int(p.R)-int((*compressedImg).ImageData[i][j].R)), 2.0)
			g := math.Pow(float64(int(p.G)-int((*compressedImg).ImageData[i][j].G)), 2.0)
			b := math.Pow(float64(int(p.B)-int((*compressedImg).ImageData[i][j].B)), 2.0)
			mse += r + g + b
			mseR += r
			mseG += g
			mseB += b
		}
	}

	mse /= float64(3 * (*originalImg).Width * (*originalImg).Height)
	mseR /= float64((*originalImg).Width * (*originalImg).Height)
	mseG /= float64((*originalImg).Width * (*originalImg).Height)
	mseB /= float64((*originalImg).Width * (*originalImg).Height)

	return mse, mseR, mseG, mseB
}

func SNR(img *lab5.TGAImage, MSE float64, MSE_R float64, MSE_G float64, MSE_B float64) (float64, float64, float64, float64) {
	var snr float64 = 0.0
	var snrR float64 = 0.0
	var snrG float64 = 0.0
	var snrB float64 = 0.0

	for _, row := range (*img).ImageData {
		for _, p := range row {
			r := math.Pow(float64(p.R), 2.0)
			g := math.Pow(float64(p.G), 2.0)
			b := math.Pow(float64(p.B), 2.0)
			snr += r + g + b
			snrR += r
			snrG += g
			snrB += b
		}
	}

	snr /= (3 * float64((*img).Width) * float64((*img).Height) * MSE)
	snrR /= (float64((*img).Width) * float64((*img).Height) * MSE_R)
	snrG /= (float64((*img).Width) * float64((*img).Height) * MSE_G)
	snrB /= (float64((*img).Width) * float64((*img).Height) * MSE_B)

	return snr, snrR, snrG, snrB
}
