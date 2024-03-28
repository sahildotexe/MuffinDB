package swah

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/sahildotexe/swah-db/kdtree"
)

type VectorStore struct {
	Tree *kdtree.KDTree
}

type Embedding struct {
	Text  string
	Point []float32
}

func Connect() *VectorStore {

	// Connect to the VectorStore
	// @param nil
	// @return *VectorStore

	var store *VectorStore
	if _, err := os.Stat("data.gob"); err == nil {
		deserializedKdtree, err := Deserialize("data.gob")
		if err != nil {
			panic(err)
		}
		store = deserializedKdtree
		fmt.Println("Loaded existing data from file")
	} else if os.IsNotExist(err) {
		fmt.Println("No existing data found, creating new Vector Store")
		store = NewVectorStore()
	} else {
		panic(err)
	}
	return store
}

func NewVectorStore() *VectorStore {

	// Create a new VectorStore
	// @param nil
	// @return *VectorStore

	return &VectorStore{
		Tree: kdtree.NewKDTree(),
	}
}

func (vs *VectorStore) InsertVector(text string, point []float32) {

	// Insert a vector into the store
	// @param point []float32
	// @return void

	id := uuid.New().String()

	v := kdtree.Vector{
		ID:     id,
		Values: point,
		Text:   text,
	}

	vs.Tree.Insert(v)
	if err := Serialize(vs, "data.gob"); err != nil {
		panic(err)
	}
}

func (vs *VectorStore) GetAllVectors() []kdtree.Vector {

	// Get all vectors from the store
	// @param nil
	// @return []kdtree.Vector

	return vs.Tree.GetAllVectors()
}

func (vs *VectorStore) GetKNearestNeighbors(point []float32, k int) []kdtree.HeapVector {

	// Get the nearest neighbor to a point
	// @param point []float32, k int
	// @return []kdtree.HeapVector

	total := vs.Tree.CountVectors()

	if k > total {
		k = total
	}

	target := kdtree.Vector{
		Values: point,
	}

	neighbours := vs.Tree.GetNeighbours(target)

	if len(neighbours) < k {
		return neighbours
	}

	return neighbours[:k]
}

func (vs *VectorStore) GetVector(id string) kdtree.Vector {

	// Get a vector from the store
	// @param id string
	// @return kdtree.Vector

	vec, found := vs.Tree.GetNodeByVectorID(id)

	if !found {
		return kdtree.Vector{}
	}

	return vec
}

func (vs *VectorStore) DeleteVector(id string) {

	// Delete a vector from the store
	// @param id string
	// @return void

	vs.Tree.DeleteNodeByVectorID(id)
	if err := Serialize(vs, "data.gob"); err != nil {
		panic(err)
	}
}

func (vs *VectorStore) UpdateVector(id string, point []float32) {

	// Update a vector in the store
	// @param id string, point []float32
	// @return void

	v := kdtree.Vector{
		ID:     id,
		Values: point,
	}

	vs.Tree.DeleteNodeByVectorID(id)

	vs.Tree.Insert(v)
	if err := Serialize(vs, "data.gob"); err != nil {
		panic(err)
	}
}
