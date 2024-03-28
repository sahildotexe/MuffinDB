package arka

import (
	"encoding/gob"
	"os"

	"github.com/sahildotexe/ArkaDB/kdtree"
)

func Serialize(store *VectorStore, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	gob.Register(VectorStore{})
	gob.Register(kdtree.Vector{})
	gob.Register(kdtree.Leaf{})
	gob.Register(kdtree.Internal{})
	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(store); err != nil {
		return err
	}

	return nil
}

func Deserialize(filename string) (*VectorStore, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var store VectorStore
	gob.Register(VectorStore{})
	gob.Register(kdtree.Leaf{})
	gob.Register(kdtree.Internal{})
	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&store); err != nil {
		return nil, err
	}

	return &store, nil
}
