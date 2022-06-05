package lab6

import (
	"KiKD/lab2"
	"KiKD/lab5"
	"math"
	"os"
)

const (
	low  = iota
	high = iota
)

func Encode(filenameIn string, filenameOut string, quantizerBits int) {
	var img lab5.TGAImage
	(&img).LoadImage(filenameIn)

	var length int
	entries := int(math.Pow(2.0, float64(quantizerBits)))

	red := make([]int, 0)
	green := make([]int, 0)
	blue := make([]int, 0)

	// make slices of int values for all three colour components
	for i := 0; i < img.Height; i++ {
		for j := 0; j < img.Width; j++ {
			red = append(red, int(img.ImageData[i][j].R))
			green = append(green, int(img.ImageData[i][j].G))
			blue = append(blue, int(img.ImageData[i][j].B))
		}
	}

	if len(red)%2 != 0 {
		red = append(red, 0)
		green = append(green, 0)
		blue = append(blue, 0)
	}

	length = len(red) // equal to len(green) and len(blue) since each pixel has those three values

	/* apply filters
	values for low are from range [0, 255]
	values for high are from range [-128, 127] -> +128 when writing to file
	*/
	for i := 0; i < length-1; i += 2 {
		red[i], red[i+1] = filter(red[i], red[i+1])
		green[i], green[i+1] = filter(green[i], green[i+1])
		blue[i], blue[i+1] = filter(blue[i], blue[i+1])
	}

	// calculate how many times each value repeats
	rl := make(map[int]int)
	rh := make(map[int]int)
	gl := make(map[int]int)
	gh := make(map[int]int)
	bl := make(map[int]int)
	bh := make(map[int]int)

	for i := 0; i < length-1; i += 2 {
		rl[red[i]]++
		rh[red[i+1]]++
		gl[green[i]]++
		gh[green[i+1]]++
		bl[blue[i]]++
		bh[blue[i+1]]++
	}

	// test returned values
	// fmt.Println("Map of red low sequence values")
	// fmt.Println(rl)
	// fmt.Println("Map of red high sequence values")
	// fmt.Println(rh)
	// fmt.Println("Map of green low sequence values")
	// fmt.Println(gl)
	// fmt.Println("Map of green high sequence values")
	// fmt.Println(gh)
	// fmt.Println("Map of blue low sequence values")
	// fmt.Println(bl)
	// fmt.Println("Map of blue high sequence values")
	// fmt.Println(bh)

	// nonuniform quantization
	usedRL, quantizerRL, mapRL := quantize(rl, len(red)/2, entries, low)
	usedRH, quantizerRH, mapRH := quantize(rh, len(red)/2, entries, high)
	usedGL, quantizerGL, mapGL := quantize(gl, len(green)/2, entries, low)
	usedGH, quantizerGH, mapGH := quantize(gh, len(green)/2, entries, high)
	usedBL, quantizerBL, mapBL := quantize(bl, len(blue)/2, entries, low)
	usedBH, quantizerBH, mapBH := quantize(bh, len(blue)/2, entries, high)

	// compression to desired bitrate
	redBytesL := make([]byte, 0)
	redBytesH := make([]byte, 0)
	greenBytesL := make([]byte, 0)
	greenBytesH := make([]byte, 0)
	blueBytesL := make([]byte, 0)
	blueBytesH := make([]byte, 0)

	for i := 0; i < length-1; i += 2 {
		redBytesL = append(redBytesL, byte(indexof(quantizerRL, mapRL[red[i]])))
		redBytesH = append(redBytesH, byte(indexof(quantizerRH, mapRH[red[i+1]])))
		greenBytesL = append(greenBytesL, byte(indexof(quantizerGL, mapGL[green[i]])))
		greenBytesH = append(greenBytesH, byte(indexof(quantizerGH, mapGH[green[i+1]])))
		blueBytesL = append(blueBytesL, byte(indexof(quantizerBL, mapBL[blue[i]])))
		blueBytesH = append(blueBytesH, byte(indexof(quantizerBH, mapBH[blue[i+1]])))
	}

	redBytesL = compress(redBytesL, quantizerBits)
	redBytesH = compress(redBytesH, quantizerBits)
	greenBytesL = compress(greenBytesL, quantizerBits)
	greenBytesH = compress(greenBytesH, quantizerBits)
	blueBytesL = compress(blueBytesL, quantizerBits)
	blueBytesH = compress(blueBytesH, quantizerBits)

	// write to file
	// quantization header
	quantHeader := make([]byte, 0)
	quantHeader = append(quantHeader, byte(len(img.Footer)))
	quantHeader = append(quantHeader, byte(quantizerBits))
	quantHeader = append(quantHeader, byte(usedRL))
	for i := 0; i < usedRL; i++ {
		quantHeader = append(quantHeader, byte(quantizerRL[i]))
	}
	quantHeader = append(quantHeader, byte(usedRH))
	for i := 0; i < usedRH; i++ {
		quantHeader = append(quantHeader, byte(quantizerRH[i]+128))
	}
	quantHeader = append(quantHeader, byte(usedGL))
	for i := 0; i < usedGL; i++ {
		quantHeader = append(quantHeader, byte(quantizerGL[i]))
	}
	quantHeader = append(quantHeader, byte(usedGH))
	for i := 0; i < usedGH; i++ {
		quantHeader = append(quantHeader, byte(quantizerGH[i]+128))
	}
	quantHeader = append(quantHeader, byte(usedBL))
	for i := 0; i < usedBL; i++ {
		quantHeader = append(quantHeader, byte(quantizerBL[i]))
	}
	quantHeader = append(quantHeader, byte(usedBH))
	for i := 0; i < usedBH; i++ {
		quantHeader = append(quantHeader, byte(quantizerBH[i]+128))
	}

	out, err := os.Create(filenameOut)
	lab2.Check(err)
	defer out.Close()

	_, err = out.Write(quantHeader)
	lab2.Check(err)

	_, err = out.Write(img.Header)
	lab2.Check(err)

	_, err = out.Write(img.Footer)
	lab2.Check(err)

	_, err = out.Write(redBytesL)
	lab2.Check(err)

	_, err = out.Write(redBytesH)
	lab2.Check(err)

	_, err = out.Write(greenBytesL)
	lab2.Check(err)

	_, err = out.Write(greenBytesH)
	lab2.Check(err)

	_, err = out.Write(blueBytesL)
	lab2.Check(err)

	_, err = out.Write(blueBytesH)
	lab2.Check(err)
}

func filter(x1 int, x2 int) (int, int) {
	var y int
	var z int

	y = (x1 + x2) / 2

	if x1 > x2 {
		z = (x2 - x1 - 1) / 2
	} else {
		z = (x2 - x1) / 2
	}

	return y, z
}

func quantize(data map[int]int, nValues int, entries int, seq int) (int, []int, map[int]int) {
	index := 0
	rangeValue := nValues / entries
	sum := 0
	lo := 0
	hi := 255
	quantizerDict := make([]int, entries)
	quantizerRanges := make([][]int, entries)
	dataMap := make(map[int]int)

	if len(data) <= entries {
		i := 0
		for k := range data {
			quantizerDict[i] = k
			dataMap[k] = k
			i++
		}

		return len(data), quantizerDict, dataMap
	}

	if seq == high {
		lo -= 128
		hi -= 128
	}

	for i := range quantizerDict {
		quantizerDict[i] = -1
		quantizerRanges[i] = make([]int, 0)
	}

	for k := lo; k <= hi; k++ {
		if v, ok := data[k]; ok {
			quantizerRanges[index] = append(quantizerRanges[index], k)

			if index == entries-1 {
				continue
			}

			sum += v

			if sum >= rangeValue {
				diff := abs(sum, rangeValue)
				diffLow := abs(sum-v, rangeValue)
				if diff < diffLow || sum-v == 0 {
					sum = 0
				} else {
					quantizerRanges[index] = quantizerRanges[index][:len(quantizerRanges[index])-1] // remove last added element since it exceeds desired range of values too much
					quantizerRanges[index+1] = append(quantizerRanges[index+1], k)                  // add that same element to the next subset since loop won't go over it again
					sum = v                                                                         // set sum as v as it would've been done otherwise with sum = 0; sum += v
				}
				index++
			}
		}
	}

	// weighted average to determine each quantizer range centre
	for i, r := range quantizerRanges {
		avgSum := 0.0
		avgDiv := 0.0
		for _, v := range r {
			avgSum += float64(v * data[v])
			avgDiv += float64(data[v])
		}
		if avgSum < 0 {
			quantizerDict[i] = int((avgSum)/avgDiv - 0.5)
		} else {
			quantizerDict[i] = int((avgSum)/avgDiv + 0.5)
		}
	}

	// mapping each value from data to corresponding quantizedDict entry
	for k := range data {
		for i, r := range quantizerRanges {
			for _, j := range r {
				if k == j {
					dataMap[k] = quantizerDict[i]
					break
				}
			}
		}
	}

	return index + 1, quantizerDict, dataMap
}

func compress(data []byte, bitrate int) []byte {
	compressedData := make([]byte, 0)
	stringData := ""

	for _, b := range data {
		stringData += byteToString(b, bitrate)
	}

	for {
		if len(stringData)%8 == 0 {
			break
		}

		stringData += "0"
	}

	for i := 0; i < len(stringData)-7; i += 8 {
		compressedData = append(compressedData, stringToByte(stringData[i:i+8]))
	}

	return compressedData
}
