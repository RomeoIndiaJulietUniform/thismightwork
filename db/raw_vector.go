package db

import (
	"encoding/json"

	"github.com/RomeoIndiaJulietUniform/thismightwork/db/base"
)

type RawVector struct {
	ID       string    `json:"id"`
	Vector   []float32 `json:"vector"`
	Metadata string    `json:"metadata"`
}

type RawVectorStore struct {
	*base.Badger
}

func (r *RawVectorStore) PutVector(v *RawVector) error {
	key := []byte("raw:" + v.ID)

	val, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return r.Badger.Put(key, val)
}

func (r *RawVectorStore) GetVector(id string) (*RawVector, error) {
	val, err := r.Badger.Get([]byte("raw:" + id))
	if err != nil {
		return nil, err
	}

	var v RawVector
	if err := json.Unmarshal(val, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

func (r *RawVectorStore) DeleteVector(id string) error {
	return r.Badger.Delete([]byte("raw:" + id))
}
