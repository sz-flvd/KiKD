package lab3

import "math"

func first(b []byte, _ error) []byte {
	return b
}

func Entropy(data map[byte]uint) float64 {
	entropy := 0.0
	sum := uint(0)

	for _, s := range data {
		sum += s
	}

	for _, s := range data {
		entropy += float64(s) / float64(sum) * -math.Log2(float64(s)/float64(sum))
	}

	return entropy
}

// func lab3Main() {
// 	if len(os.Args) == 4 || len(os.Args) == 5 {
// 		start := time.Now()

// 		if len(os.Args) == 5 {
// 			var param int
// 			if os.Args[2] == "-g" {
// 				param = lab3.Gamma
// 			} else if os.Args[2] == "-d" {
// 				param = lab3.Delta
// 			} else if os.Args[2] == "-o" {
// 				param = lab3.Omega
// 			} else if os.Args[2] == "-f" {
// 				param = lab3.Fib
// 			} else {
// 				fmt.Println("Incorrent argument - must be [-g/-d/-o/-f]")
// 				return
// 			}

// 			if os.Args[1] == "-enc" {
// 				lab3.Encode(os.Args[3], os.Args[4], param)
// 			} else if os.Args[1] == "-dec" {
// 				lab3.Decode(os.Args[3], os.Args[4], param)
// 			}
// 		} else {
// 			if os.Args[1] == "-enc" {
// 				lab3.Encode(os.Args[2], os.Args[3], lab3.Omega)
// 			} else if os.Args[1] == "-dec" {
// 				lab3.Decode(os.Args[2], os.Args[3], lab3.Omega)
// 			}
// 		}

// 		elapsed := time.Since(start)

// 		fmt.Printf("-----Elapsed time: %s-----\n", elapsed)
// 	} else {
// 		fmt.Println("Usage:\tgo run main.go -enc [-g/-d/-o/-f] file_to_encode output_file\n\tor")
// 		fmt.Println("\tgo run main.go -dec [-g/-d/-o/-f] file_to_decode output_file")
// 	}
// }
