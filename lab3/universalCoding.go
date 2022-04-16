package lab3

import (
	"fmt"
	"strconv"
)

func eliasGamma(num uint64) string {
	code := strconv.FormatUint(num, 2)
	n := len(code)

	for i := 0; i < n-1; i++ {
		code = "0" + code
	}

	return code
}

func eliasGammaDecode(code string) (uint64, uint64) {
	var n uint64
	counter := 0

	for {
		if code[counter] == '0' {
			counter++
		} else {
			break
		}
	}

	n, _ = strconv.ParseUint(code[counter:2*counter+1], 2, 64)

	return n, uint64(2*counter + 1)
}

func eliasDelta(num uint64) string {
	code := strconv.FormatUint(num, 2)
	n := len(code)
	prefix := eliasGamma(uint64(n))
	code = prefix + code[1:]

	return code
}

func eliasDeltaDecode(code string) (uint64, uint64) {
	var x uint64
	var n uint64
	var k uint64
	counter := 0

	for {
		if code[counter] == '0' {
			counter++
		} else {
			break
		}
	}

	k = uint64(counter + 1)
	n, _ = strconv.ParseUint(code[counter:counter+int(k)], 2, 64)
	x, _ = strconv.ParseUint("1"+code[counter+int(k):counter+int(k+n-1)], 2, 64)

	return x, uint64(counter)
}

func eliasOmega(num uint64) string {
	code := "0"
	k := num

	for {
		if k <= 1 {
			break
		}

		code = strconv.FormatUint(k, 2) + code

		k = uint64(len(strconv.FormatUint(k, 2)))
	}

	return code
}

func eliasOmegaDecode(code string) (uint64, uint64) {
	n := uint64(1)
	counter := 0

	for {
		if code[counter] == '0' {
			break
		}

		counter = counter + int(n) + 1
		n, _ = strconv.ParseUint(code[counter-int(n)-1:counter], 2, 64)
	}

	return n, uint64(counter + 1)
}

func fibonacci(num uint64) string {
	code := "1"
	f, a := fibonacciNumber(num)
	coeffs := make([]int, a+1)

	for i := range coeffs {
		coeffs[i] = 0
	}

	coeffs[a] = 1
	num -= uint64(f)

	for {
		if num == 0 {
			break
		}

		f, a = fibonacciNumber(num)
		coeffs[a] = 1
		num -= uint64(f)
	}

	for i := len(coeffs) - 1; i >= 0; i-- {
		code = fmt.Sprint(coeffs[i]) + code
	}

	return code
}

func fibonacciDecode(code string) (uint64, uint64) {
	x := uint64(0)
	f := make([]int, 2)
	f[0] = 1
	f[1] = 1
	i := 0

	for {
		if code[i] == '1' && code[i-1] == '1' {
			break
		}

		if code[i] == '1' {
			x += uint64(f[i+1])
		}

		f = append(f, f[i]+f[i+1])
		i++
	}

	return x, uint64(i + 1)
}

func fibonacciNumber(k uint64) (int, int) {
	f := make([]int, 2)
	f[0] = 1
	f[1] = 1
	i := 1

	for {
		if uint64(f[i]) > k {
			break
		}

		f = append(f, f[i]+f[i-1])
		i++
	}

	i--

	return f[i], i
}
