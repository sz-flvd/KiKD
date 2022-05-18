package main

import (
	"KiKD/lab2"
	"KiKD/lab5"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func lab5Main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go img_in.tga img_out.tga colour_depth")
		return
	} else {
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
		if k < 0 || k > 24 {
			fmt.Println("Colour depth should be between 0 and 24")
			return
		}
		start := time.Now()
		outImg := lab5.Quantize(&img, k)
		elapsed := time.Since(start)
		(&outImg).SaveImage(out)

		mse = lab5.MeanSquaredError(&img, &outImg)
		snr = lab5.SignalToNoiseRatio(&img, mse)

		fmt.Printf("Elapsed time: %s\n", elapsed)
		fmt.Printf("Mean squared error: %f\n", mse)
		fmt.Printf("Signal-to-noise ratio: %f [%f dB]\n", snr, 10*math.Log10(snr))
	}
}

func main() {
	lab5Main()
}
