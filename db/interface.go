package db

type KVStore interface {
	Put(key, value []byte) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
	Close()
}
