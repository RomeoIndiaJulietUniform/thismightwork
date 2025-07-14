package db

import (
	"log"

	badger "github.com/dgraph-io/badger"
)

type Badger struct {
	db *badger.DB
}

var _ KVStore = (*Badger)(nil)

func BadgerDB(path string) *Badger {
	opts := badger.DefaultOptions(path).WithLogger(nil)
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatalf("Failed to open BadgerDB: %v", err)
	}
	return &Badger{db: db}
}

func (b *Badger) Put(key, value []byte) error {
	return b.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
}

func (b *Badger) Get(key []byte) ([]byte, error) {
	var valCopy []byte
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		valCopy = val
		return nil
	})
	return valCopy, err
}

func (b *Badger) Delete(key []byte) error {
	return b.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}

func (b *Badger) Close() {
	if err := b.db.Close(); err != nil {
		log.Printf("Error closing BadgerDB: %v", err)
	}
}
