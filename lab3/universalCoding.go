package lab3

import (
	"strconv"
)

func eliasGamma(num int) string {
	code := strconv.FormatInt(int64(num), 2)
	n := len(code)

	for i := 0; i < n-1; i++ {
		code = "0" + code
	}

	return code
}

func eliasGammaDecode(code string) (int, int) {
	var n int64
	counter := 0

	for {
		if counter >= len(code)-1 {
			return 0, 0
		}

		if code[counter] == '0' {
			counter++
		} else {
			break
		}
	}

	if len(code) <= 2*counter {
		return 0, 0
	}

	n, _ = strconv.ParseInt(code[counter:2*counter+1], 2, 64)

	return int(n), 2*counter + 1
}

func eliasDelta(num int) string {
	code := strconv.FormatInt(int64(num), 2)
	n := len(code)
	prefix := eliasGamma(n)
	code = prefix + code[1:]

	return code
}

func eliasDeltaDecode(code string) (int, int) {
	var x int64
	var n int64
	var k int
	counter := 0

	for {
		if counter >= len(code)-1 {
			return 0, 0
		}

		if code[counter] == '0' {
			counter++
		} else {
			break
		}
	}

	k = counter + 1

	if len(code) <= 2*counter {
		return 0, 0
	}

	n, _ = strconv.ParseInt(code[counter:counter+int(k)], 2, 64)

	if len(code) <= 2*counter+int(n) {
		return 0, 0
	}

	x, _ = strconv.ParseInt("1"+code[counter+k:2*counter+int(n)], 2, 64)

	return int(x), 2*counter + int(n)
}

func eliasOmega(num int) string {
	code := "0"
	k := num

	for {
		if k <= 1 {
			break
		}

		code = strconv.FormatInt(int64(k), 2) + code

		k = len(strconv.FormatInt(int64(k), 2)) - 1
	}

	return code
}

func eliasOmegaDecode(code string) (int, int) {
	n := int64(1)
	counter := 0

	if len(code) == 0 {
		return 0, 0
	}

	for {
		if code[counter] == '0' {
			break
		}

		counter = counter + int(n) + 1

		if counter >= len(code)-1 {
			return 0, 0
		}

		n, _ = strconv.ParseInt(code[counter-int(n)-1:counter], 2, 64)
	}

	return int(n), counter + 1
}

func fibonacci(num int) string {
	code := ""
	f := fibonacciNumber(num)

	for i := len(f) - 1; i >= 0; i-- {
		if f[i] <= num {
			code = "1" + code
			num -= f[i]
		} else {
			code = "0" + code
		}
	}

	code += "1"

	return code
}

func fibonacciDecode(code string) (int, int) {
	x := 0
	f := make([]int, 2)
	f[0] = 1
	f[1] = 2
	i := 0

	for {
		if i >= len(code)-1 {
			return 0, 0
		}

		if code[i] == '1' {
			x += f[i]

			if code[i+1] == '1' {
				break
			}
		}

		f = append(f, f[i]+f[i+1])
		i++
	}

	return x, i + 2
}

func fibonacciNumber(k int) []int {
	f := make([]int, 2)
	f[0] = 1
	f[1] = 2
	i := 0

	for {
		if f[i] > k {
			break
		}

		f = append(f, f[i]+f[i+1])
		i++
	}

	return f[:len(f)-2]
}
