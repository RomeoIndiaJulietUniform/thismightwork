package algo

import (
	"testing"

	"github.com/RomeoIndiaJulietUniform/thismightwork/algo"
)

func TestKMeans(t *testing.T) {
	points := [][]float64{
		{1.0, 2.0},
		{1.1, 2.1},
		{0.9, 1.9},
		{8.0, 8.0},
		{8.1, 8.1},
		{7.9, 7.9},
	}

	k := 2
	iterations := 10

	centroids := algo.KMeans(points, k, iterations)

	if len(centroids) != k {
		t.Fatalf("Expected %d centroids, got %d", k, len(centroids))
	}

	for _, c := range centroids {
		if len(c) != 2 {
			t.Errorf("Expected 2D centroids, got %v", c)
		}
		if c[0] < 0 || c[1] < 0 {
			t.Errorf("Unexpected negative coordinate in centroid: %v", c)
		}
	}
}

func TestAssignToClusters(t *testing.T) {
	points := [][]float64{
		{0, 0},
		{1, 1},
		{10, 10},
		{11, 11},
	}
	centroids := [][]float64{
		{0, 0},
		{10, 10},
	}

	clusters := algo.AssignToClusters(points, centroids)

	if len(clusters) != 2 {
		t.Fatalf("Expected 2 clusters, got %d", len(clusters))
	}

	if clusters[0][0] != 0 && clusters[0][1] != 1 {
		t.Errorf("Expected point 0 and 1 in cluster 0, got %v", clusters[0])
	}
}
