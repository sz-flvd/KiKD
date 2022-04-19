package main

import (
	"KiKD/lab3"
	"fmt"
	"os"
	"time"
)

// func lab2Main() {
// 	if len(os.Args) == 4 {
// 		start := time.Now()

// 		if os.Args[1] == "-e" {
// 			lab2.EncodeFile(os.Args[2], os.Args[3])
// 		} else if os.Args[1] == "-d" {
// 			lab2.DecodeFile(os.Args[2], os.Args[3])
// 		}

// 		elapsed := time.Since(start)

// 		fmt.Printf("-----Elapsed time: %s-----\n", elapsed)
// 	} else {
// 		fmt.Println("Usage:\tgo run main.go -e file_to_encode output_file\n\tor")
// 		fmt.Println("\tgo run main.go -d file_to_decode output_file")
// 	}
// }

func lab3Main() {
	if len(os.Args) == 4 || len(os.Args) == 5 {
		start := time.Now()

		if len(os.Args) == 5 {
			var param int
			if os.Args[2] == "-g" {
				param = lab3.Gamma
			} else if os.Args[2] == "-d" {
				param = lab3.Delta
			} else if os.Args[2] == "-o" {
				param = lab3.Omega
			} else if os.Args[2] == "-f" {
				param = lab3.Fib
			} else {
				fmt.Println("Incorrent argument - must be [-g/-d/-o/-f]")
				return
			}

			if os.Args[1] == "-enc" {
				lab3.Encode(os.Args[3], os.Args[4], param)
			} else if os.Args[1] == "-dec" {
				lab3.Decode(os.Args[3], os.Args[4], param)
			}
		} else {
			if os.Args[1] == "-enc" {
				lab3.Encode(os.Args[2], os.Args[3], lab3.Omega)
			} else if os.Args[1] == "-dec" {
				lab3.Decode(os.Args[2], os.Args[3], lab3.Omega)
			}
		}

		elapsed := time.Since(start)

		fmt.Printf("-----Elapsed time: %s-----\n", elapsed)
	} else {
		fmt.Println("Usage:\tgo run main.go -enc [-g/-d/-o/-f] file_to_encode output_file\n\tor")
		fmt.Println("\tgo run main.go -dec [-g/-d/-o/-f] file_to_decode output_file")
	}
}

func main() {
	lab3Main()
}
