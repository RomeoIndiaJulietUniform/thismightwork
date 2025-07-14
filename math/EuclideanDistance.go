package math

import (
	"log"
	"math"
)

func EuclideanDistance(p1, p2 []float64) float64 {
	if len(p1) != len(p2) {
		log.Fatal("Vector dimension mismatch")
	}

	sum := 0.0
	for i := range p1 {
		diff := p1[i] - p2[i]
		sum += diff * diff
	}

	return math.Sqrt(sum)
}
