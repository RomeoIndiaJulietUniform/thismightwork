package db

import "log"

func newSession(backend, data string) KVStore {
	switch backend {
	case "badgerdb":
		return BadgerDB(data)
	default:
		log.Fatal("No DB Selected")
		return nil
	}
}
