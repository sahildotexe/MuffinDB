package main

import (
	"GoVectorDB/store"
	"fmt"
)

func main() {
	store := store.NewVectorStore()
	store.InsertVector("Mango", []float32{1, 2, 3})
	store.InsertVector("Apple", []float32{4, 5, 6})
	store.InsertVector("Orange", []float32{7, 8, 9})

	vecs := store.GetAllVectors()
	for _, v := range vecs {
		fmt.Println(v)
	}

	tar := []float32{1, 2, 3}
	neighbours := store.GetKNearestNeighbors(tar, 5)

	for _, v := range neighbours {
		fmt.Println(v)
	}
}
