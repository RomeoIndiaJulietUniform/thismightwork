package db

import (
	"encoding/json"

	"github.com/RomeoIndiaJulietUniform/thismightwork/db/base"
)

type HNSWVector struct {
	ID         string   `json:"id"`
	Level      int      `json:"level"`
	Neighbours []string `json:"neighbours"`
}

type HNSWStore struct {
	*base.Badger
}

func (h *HNSWStore) PutVector(node *HNSWVector) error {
	key := []byte("hnsw:" + node.ID)
	val, err := json.Marshal(node)
	if err != nil {
		return err
	}
	return h.Badger.Put(key, val)
}

func (h *HNSWStore) GetVector(id string) (*HNSWVector, error) {
	val, err := h.Badger.Get([]byte("hnsw:" + id))
	if err != nil {
		return nil, err
	}

	var node HNSWVector
	if err := json.Unmarshal(val, &node); err != nil {
		return nil, err
	}
	return &node, nil
}
