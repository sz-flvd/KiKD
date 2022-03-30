package main

import (
	"KiKD/lab2"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 4 {
		if os.Args[1] == "-e" {
			lab2.EncodeFile(os.Args[2], os.Args[3])
		} else if os.Args[1] == "-d" {
			lab2.DecodeFile(os.Args[2], os.Args[3])
		}
	} else {
		fmt.Println("Usage:\tgo run main.go -e file_to_encode output_file\n\tor")
		fmt.Println("\tgo run main.go -d file_to_decode output_file")
	}
}
