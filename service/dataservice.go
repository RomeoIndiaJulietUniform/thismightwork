package service

import (
	"context"
	"log"

	"github.com/RomeoIndiaJulietUniform/thismightwork/api/pb"
	"github.com/RomeoIndiaJulietUniform/thismightwork/db"
	"github.com/RomeoIndiaJulietUniform/thismightwork/db/base"
)

type DataService struct {
	pb.UnimplementedSeirraRomeoServer
	rawStore  *db.RawVectorStore
	hnswStore *db.HNSWStore
}

func NewDataService(dbPath string) *DataService {
	badger := base.BadgerDB(dbPath)
	return &DataService{
		rawStore:  &db.RawVectorStore{Badger: badger},
		hnswStore: &db.HNSWStore{Badger: badger},
	}
}

func (s *DataService) Insert(ctx context.Context, req *pb.InsertRequest) (*pb.InsertResponse, error) {
	raw := &db.RawVector{
		ID:       req.Vector.Id,
		Vector:   req.Vector.Values,
		Metadata: req.Vector.Metadata["description"],
	}
	if err := s.rawStore.PutVector(raw); err != nil {
		log.Println("failed to insert raw vector:", err)
		return &pb.InsertResponse{Success: false, Message: "raw insert failed"}, err
	}

	hnsw := &db.HNSWVector{
		ID:         req.Vector.Id,
		Level:      0,
		Neighbours: []string{},
	}
	if err := s.hnswStore.PutVector(hnsw); err != nil {
		log.Println("failed to insert hnsw vector:", err)
		return &pb.InsertResponse{Success: false, Message: "hnsw insert failed"}, err
	}

	return &pb.InsertResponse{Success: true, Message: "inserted successfully"}, nil
}
