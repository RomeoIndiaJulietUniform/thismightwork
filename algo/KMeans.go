package algo

import (
	stdmath "math"
	"math/rand"
	"time"

	"github.com/RomeoIndiaJulietUniform/thismightwork/math"
)

func KMeans(points [][]float64, k int, itr int) [][]float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	centroids := kMeansPlusPlusInit(points, k, r)

	for range itr {
		clusters := AssignToClusters(points, centroids)

		for i := range clusters {
			if len(clusters[i]) == 0 {
				continue
			}

			clusterPoints := make([][]float64, 0, len(clusters[i]))
			for _, idx := range clusters[i] {
				clusterPoints = append(clusterPoints, points[idx])
			}

			centroids[i] = math.Centroid(clusterPoints)
		}
	}

	return centroids
}

func AssignToClusters(points [][]float64, centroids [][]float64) [][]int {
	clusters := make([][]int, len(centroids))

	for i := range points {
		minDist := stdmath.MaxFloat64
		bestCluster := 0

		for j := range centroids {
			dist := math.EuclideanDistance(points[i], centroids[j])
			if dist < minDist {
				minDist = dist
				bestCluster = j
			}
		}

		clusters[bestCluster] = append(clusters[bestCluster], i)
	}

	return clusters
}

func kMeansPlusPlusInit(points [][]float64, k int, r *rand.Rand) [][]float64 {
	n := len(points)
	centroids := make([][]float64, 0, k)

	first := points[r.Intn(n)]
	centroids = append(centroids, first)

	for len(centroids) < k {
		distances := make([]float64, n)
		totalDist := 0.0

		for i, p := range points {
			minDist := stdmath.MaxFloat64
			for _, c := range centroids {
				d := math.EuclideanDistance(p, c)
				if d < minDist {
					minDist = d
				}
			}
			distances[i] = minDist * minDist
			totalDist += distances[i]
		}

		randVal := r.Float64() * totalDist
		sum := 0.0
		for i, d := range distances {
			sum += d
			if randVal <= sum {
				centroids = append(centroids, points[i])
				break
			}
		}
	}

	return centroids
}
