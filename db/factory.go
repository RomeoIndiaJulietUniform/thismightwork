package db

import "log"

func newSession(backend, data string) KVStore {
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
