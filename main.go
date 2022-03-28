package main

import (
	"KiKD/lab2"
	"fmt"
	"os"
)

func main() {
	if os.Args[1] == "-e" {
		lab2.EncodeFile(os.Args[2], os.Args[3])
	} else if os.Args[1] == "-d" {
		lab2.DecodeFile(os.Args[2], os.Args[3])
	} else {
		fmt.Println("Usage: go run main.go -e file_to_encode or")
		fmt.Println("go run main.go -d file_to_decode")
	}
}
