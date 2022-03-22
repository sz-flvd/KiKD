package lab2

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadAndEncode(filename string) {
	file, err := os.Open(filename)
	defer file.Close()
	check(err)

	br := bufio.NewReader(file)

	b, e := br.ReadByte()

	check(e)
	fmt.Println(b)

	// for b,e := br.ReadByte() {
	// 	if e != nil && !errors.Is(e, io.EOF) {
	// 		fmt.Println(e)
	// 		break
	// 	}

	// 	fmt.Println(b)

	// 	if e != nil {
	// 		break
	// 	}
	// }
}

func ReadAndDecode(filename string) {

}
