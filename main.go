package main

import (
	"KiKD/lab4"
	"fmt"
	"os"
)

func lab4Main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go file.tga")
		return
	}

	_ = lab4.LoadTGAFile(os.Args[1])
	//predicted := lab4.Predict(bytes, 0)

	//fmt.Println(predicted)
}

func main() {
	lab4Main()
}
