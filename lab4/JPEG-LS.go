package lab4

func Predict(bytes *[][]byte, option int) [][]int {
	height := len(*bytes)
	width := len((*bytes)[0])
	predicted := make([][]int, height)

	for i := 0; i < height; i++ {
		predicted[i] = make([]int, width)
	}

	if option < 1 || option > 8 {
		return nil
	}

	switch option {
	case 1:
		predict1(bytes, &predicted)
	case 2:
		predict2(bytes, &predicted)
	case 3:
		predict3(bytes, &predicted)
	case 4:
		predict4(bytes, &predicted)
	case 5:
		predict5(bytes, &predicted)
	case 6:
		predict6(bytes, &predicted)
	case 7:
		predict7(bytes, &predicted)
	case 8:
		predictNew(bytes, &predicted)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			predicted[i][j] = int((*bytes)[i][j]) - predicted[i][j]

			if predicted[i][j] < 0 {
				predicted[i][j] += 256
			}

			predicted[i][j] %= 256
		}
	}

	return predicted
}

func predict1(bytes *[][]byte, predicted *[][]int) {
	for i := range *predicted {
		for j := range (*predicted)[i] {
			if j < 3 {
				(*predicted)[i][j] = 0
			} else {
				(*predicted)[i][j] = int((*bytes)[i][j-3])
			}
		}
	}
}

func predict2(bytes *[][]byte, predicted *[][]int) {
	for i := range *predicted {
		for j := range (*predicted)[i] {
			if i == 0 {
				(*predicted)[i][j] = 0
			} else {
				(*predicted)[i][j] = int((*bytes)[i-1][j])
			}
		}
	}
}

func predict3(bytes *[][]byte, predicted *[][]int) {
	for i := range *predicted {
		for j := range (*predicted)[i] {
			if i == 0 || j < 3 {
				(*predicted)[i][j] = 0
			} else {
				(*predicted)[i][j] = int((*bytes)[i-1][j-3])
			}
		}
	}
}

func predict4(bytes *[][]byte, predicted *[][]int) {
	for i := range *predicted {
		for j := range (*predicted)[i] {
			if i == 0 {
				(*predicted)[i][j] = int((*bytes)[i][j-3])
			} else if j < 3 {
				(*predicted)[i][j] = int((*bytes)[i-1][j])
			} else {
				(*predicted)[i][j] = int((*bytes)[i-1][j]) + int((*bytes)[i][j-3]) - int((*bytes)[i-1][j-3])
			}
		}
	}
}

func predict5(bytes *[][]byte, predicted *[][]int) {
	for i := range *predicted {
		for j := range (*predicted)[i] {
			if i == 0 {
				(*predicted)[i][j] = int((*bytes)[i][j-3]) / 2
			} else if j < 3 {
				(*predicted)[i][j] = int((*bytes)[i-1][j])
			} else {
				(*predicted)[i][j] = int((*bytes)[i-1][j]) + (int((*bytes)[i][j-3])-int((*bytes)[i-1][j-3]))/2
			}
		}
	}
}

func predict6(bytes *[][]byte, predicted *[][]int) {
	for i := range *predicted {
		for j := range (*predicted)[i] {
			if i == 0 {
				(*predicted)[i][j] = int((*bytes)[i][j-3])
			} else if j < 3 {
				(*predicted)[i][j] = int((*bytes)[i-1][j]) / 2
			} else {
				(*predicted)[i][j] = int((*bytes)[i][j-3]) + (int((*bytes)[i-1][j])-int((*bytes)[i-1][j-3]))/2
			}
		}
	}
}

func predict7(bytes *[][]byte, predicted *[][]int) {
	for i := range *predicted {
		for j := range (*predicted)[i] {
			if i == 0 {
				(*predicted)[i][j] = int((*bytes)[i][j-3]) / 2
			} else if j < 3 {
				(*predicted)[i][j] = int((*bytes)[i-1][j]) / 2
			} else {
				(*predicted)[i][j] = (int((*bytes)[i-1][j]) + int((*bytes)[i][j-3])) / 2
			}
		}
	}
}

func predictNew(bytes *[][]byte, predicted *[][]int) {

}
