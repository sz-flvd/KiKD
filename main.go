package main

import (
	"KiKD/lab2"
	"KiKD/lab5"
	"fmt"
	"math"
	"os"
	"strconv"
)

func lab5Main() {
	var img lab5.TGAImage
	var out string
	var k int
	var e error
	var mse float64
	var snr float64

	(&img).LoadImage(os.Args[1])
	out = os.Args[2]
	k, e = strconv.Atoi(os.Args[3])
	lab2.Check(e)
	outImg := lab5.Quantize(&img, k)
	(&outImg).SaveImage(out)

	mse = lab5.MeanSquaredError(&img, &outImg)
	snr = lab5.SignalToNoiseRatio(&img, mse)

	fmt.Printf("Mean squared error: %f\n", mse)
	fmt.Printf("Signal-to-noise ratio: %f [%f dB]\n", snr, 10*math.Log10(snr))
}

func main() {
	lab5Main()
}
