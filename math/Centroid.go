package math

import "log"

func Centroid(points [][]float64) []float64 {
	dim := len(points[0])
	result := make([]float64, dim)

	for _, point := range points {
		if len(point) != dim {
			log.Fatal("Dimension Mismatch")
		}

		for i := range dim {
			result[i] += point[i]
		}
	}

	for i := range dim {
		result[i] /= float64(len(points))
	}

	return result
}
