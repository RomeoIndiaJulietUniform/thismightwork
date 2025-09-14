package service

import (
	"github.com/RomeoIndiaJulietUniform/thismightwork/db"
	"github.com/RomeoIndiaJulietUniform/thismightwork/index"
)

type HNSWService struct {
	rawStore *db.RawVectorStore
	hnsw     *index.HNSW
}

func NewHNSWService(rawStore *db.RawVectorStore, M, EF, EFConstruction int) *HNSWService {
	return &HNSWService{
		rawStore: rawStore,
		hnsw:     index.NewHNSW(M, EF, EFConstruction),
	}
}

func (s *HNSWService) AddToIndex(id string, vector []float64) *index.Node {
	return s.hnsw.AddNode(id, vector)
}

func (s *HNSWService) GetNeighborIDs(node *index.Node) []string {
	ids := []string{}
	for _, neighbor := range node.Neighbors[0] {
		ids = append(ids, neighbor.ID)
	}
	return ids
}

func (s *HNSWService) SearchKNN(vector []float64, k int) [][]float64 {
	return s.hnsw.SearchKNN(vector, k)
}
