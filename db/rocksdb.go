package db

import (
	"github.com/linxGnu/grocksdb"
)

var _ KVStore = (*rocks)(nil)

type rocks struct {
	db *grocksdb.DB
}

func RocksDB(path string) (*rocks, error) {
	opts := grocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)

	db, err := grocksdb.OpenDb(opts, path)
	if err != nil {
		return nil, err
	}
	return &rocks{db: db}, nil
}

func (r *rocks) Put(key, value []byte) error {
	wo := grocksdb.NewDefaultWriteOptions()
	defer wo.Destroy()
	return r.db.Put(wo, key, value)
}

func (r *rocks) Get(key []byte) ([]byte, error) {
	ro := grocksdb.NewDefaultReadOptions()
	defer ro.Destroy()

	val, err := r.db.Get(ro, key)
	if err != nil {
		return nil, err
	}
	defer val.Free()

	data := make([]byte, len(val.Data()))
	copy(data, val.Data())
	return data, nil
}

func (r *rocks) Delete(key []byte) error {
	wo := grocksdb.NewDefaultWriteOptions()
	defer wo.Destroy()
	return r.db.Delete(wo, key)
}

func (r *rocks) Close() {
	r.db.Close()
}
