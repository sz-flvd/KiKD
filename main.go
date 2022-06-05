package main

import (
	"KiKD/lab2"
	"KiKD/lab6"
	"KiKD/lab7"
	"fmt"
	"os"
	"strconv"
	"time"
)

func lab6Main() {
	if len(os.Args) == 3 {
		lab6.ShowMSESNR(os.Args[1], os.Args[2])
	} else if len(os.Args) == 5 {
		k, e := strconv.Atoi(os.Args[4])
		lab2.Check(e)

		if k < 1 || k > 7 {
			fmt.Println("Number of quantizer bits must be between 1 and 7")
			return
		}

		fmt.Println("Beginning encoding image")
		start := time.Now()
		lab6.Encode(os.Args[1], os.Args[2], k)
		elapsed := time.Since(start)
		fmt.Printf("Encoding - elapsed time: %s\n", elapsed)
		fmt.Println("Beginning decoding")
		start = time.Now()
		lab6.Decode(os.Args[2], os.Args[3])
		elapsed = time.Since(start)
		fmt.Printf("Deconding - elapsed time: %s\n", elapsed)
	} else {
		fmt.Println("Usage: \t go run main.go file_in file_out quantizer_bits")
		fmt.Println("or \t go run main.go original_img compressed_img")
		return
	}
}

func lab7Main() {
	if len(os.Args) == 5 && os.Args[1] == "szum" {
		p, err := strconv.ParseFloat(os.Args[2], 64)
		lab2.Check(err)
		lab7.GenerateNoise(p, os.Args[3], os.Args[4])
	} else if len(os.Args) == 4 && (os.Args[1] == "koder" || os.Args[1] == "dekoder" || os.Args[1] == "sprawdz") {
		if os.Args[1] == "koder" {
			lab7.Encode(os.Args[2], os.Args[3])
		} else if os.Args[1] == "dekoder" {
			lab7.Decode(os.Args[2], os.Args[3])
		} else { // os.Args[1] = "sprawdz"
			lab7.Check(os.Args[2], os.Args[3])
		}
	} else {
		fmt.Println("Usage:")
		fmt.Println("go run main.go koder in out")
		fmt.Println("go run main.go szum p in out")
		fmt.Println("go run main.go dekoder in out")
		fmt.Println("go run main.go sprawdz in1 in2")
		return
	}
}

func main() {
	lab6Main()
	lab7Main()
}
