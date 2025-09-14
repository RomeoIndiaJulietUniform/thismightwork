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
	rawStore    *db.RawVectorStore
	hnswStore   *db.HNSWStore
	hnswService *HNSWService
}

func NewDataService(dbPath string, M, EF, EFConstruction int) *DataService {
	badger := base.BadgerDB(dbPath)

	rawStore := &db.RawVectorStore{Badger: badger}
	hnswService := NewHNSWService(rawStore, M, EF, EFConstruction)

	return &DataService{
		rawStore:    rawStore,
		hnswStore:   &db.HNSWStore{Badger: badger},
		hnswService: hnswService,
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

	vector64 := make([]float64, len(req.Vector.Values))
	for i, v := range req.Vector.Values {
		vector64[i] = float64(v)
	}

	node := s.hnswService.AddToIndex(req.Vector.Id, vector64)

	hnsw := &db.HNSWVector{
		ID:         req.Vector.Id,
		Level:      node.Level,
		Neighbours: s.hnswService.GetNeighborIDs(node),
	}

	if err := s.hnswStore.PutVector(hnsw); err != nil {
		log.Println("failed to insert hnsw vector:", err)
		return &pb.InsertResponse{Success: false, Message: "hnsw insert failed"}, err
	}

	return &pb.InsertResponse{Success: true, Message: "inserted successfully"}, nil
}
