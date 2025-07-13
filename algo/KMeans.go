package algo

import (
	stdmath "math"
	"math/rand"
	"time"

	"github.com/RomeoIndiaJulietUniform/thismightwork/math"
)

func KMeans(points [][]float64, k int, itr int) [][]float64 {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	centroids := make([][]float64, k)

	for i := range k {
		centroids[i] = points[r.Intn(len(points))]
	}

	for range itr {
		clusters := assignToClusters(points, centroids)

		for i := range clusters {

			if len(clusters[i]) == 0 {
				continue
			}
			newCluster := [][]float64{}

			for _, idx := range clusters[i] {
				newCluster = append(newCluster, points[idx])
			}

			centroids[i] = math.Centroid(newCluster)

		}
	}

	return centroids
}

func assignToClusters(points [][]float64, centroids [][]float64) [][]int {

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
