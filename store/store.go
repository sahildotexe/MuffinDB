package store

import (
	"GoVectorDB/kdtree"

	"github.com/google/uuid"
)

type VectorStore struct {
	Tree *kdtree.KDTree
}

func NewVectorStore() *VectorStore {

	// Create a new VectorStore
	// @param nil
	// @return *VectorStore

	return &VectorStore{
		Tree: kdtree.NewKDTree(),
	}
}

func (vs *VectorStore) InsertVector(point []float32) {

	// Insert a vector into the store
	// @param point []float32
	// @return void

	id := uuid.New().String()

	v := kdtree.Vector{
		ID:     id,
		Values: point,
	}

	vs.Tree.Insert(v)
}

func (vs *VectorStore) GetAllVectors() []kdtree.Vector {

	// Get all vectors from the store
	// @param nil
	// @return []kdtree.Vector

	return vs.Tree.GetAllVectors()
}
