package lab5

import (
	"math"
)

func Quantize(img *TGAImage, colourCount int) TGAImage {
	var dist int
	var minDist int
	var closestEntry pixel
	imgData := (*img).FlattenImageData()

	out := img.CopyImage()

	dict := make([]pixel, 1)
	dict[0] = AveragePixel(imgData)

	LindeBuzoGray(&dict, &imgData, int(math.Pow(2.0, float64(colourCount))))

	for i, col := range (*img).imageData {
		for j, px := range col {
			minDist = math.MaxInt32
			for _, v := range dict {
				dist = TaxicabDistance(px, v)
				if dist < minDist {
					minDist = dist
					closestEntry = v
				}
			}
			out.imageData[i][j] = closestEntry
		}
	}

	return out
}

func LindeBuzoGray(dict *[]pixel, pixels *[]pixel, dictSize int) {
	var dist int
	var minDist int
	var minIndex int
	tempDict := make([]pixel, 0)
	quantizationAreas := make([][]pixel, 0)

	for {
		if len(*dict) == dictSize {
			break
		}

		for _, p := range *dict {
			tempDict = append(tempDict, p.AddPerturbationVector())
		}
		*dict = append(*dict, tempDict...)
		tempDict = tempDict[:0]

		for range *dict {
			quantizationAreas = append(quantizationAreas, make([]pixel, 0))
		}

		for _, p := range *pixels {
			minDist = math.MaxInt32
			for j, y := range *dict {
				dist = TaxicabDistance(p, y)
				if dist < minDist {
					minDist = dist
					minIndex = j
				}
			}
			quantizationAreas[minIndex] = append(quantizationAreas[minIndex], p)
		}

		*dict = (*dict)[:0]
		for _, v := range quantizationAreas {
			*dict = append(*dict, AveragePixel(v))
		}
		quantizationAreas = quantizationAreas[:0]
	}
}

func MeanSquaredError(in *TGAImage, out *TGAImage) float64 {
	mse := 0.0

	for i := range (*in).imageData {
		for j, p := range (*in).imageData[i] {
			mse += math.Pow(float64(TaxicabDistance(p, (*out).imageData[i][j])), 2)
		}
	}

	mse /= float64(((*in).width * (*in).height))
	return mse
}

func SignalToNoiseRatio(img *TGAImage, mse float64) float64 {
	snr := 0.0

	zeroPixel := pixel{
		r: 0,
		g: 0,
		b: 0}

	for i := range (*img).imageData {
		for j := range (*img).imageData[i] {
			snr += math.Pow(float64(TaxicabDistance((*img).imageData[i][j], zeroPixel)), 2)
		}
	}

	snr /= (float64((*img).width) * float64((*img).height) * mse)

	return snr
}
