package db

import (
	"fmt"
	"log"
)

func NewSession(backend, data string) KVStore {
	fmt.Printf("Any Station: %v %v", backend, data)
	switch backend {
	case "rocksdb":
		db, err := RocksDB(data)
		if err != nil {
			log.Fatalf("Failed to create RocksDB: %v", err)
		}
		return db
	case "badgerdb":
		return BadgerDB(data)
	default:
		log.Fatal("No DB Selected")
		return nil
	}
}
