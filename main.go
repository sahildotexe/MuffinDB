package main

import (
	"GoVectorDB/store"
	"fmt"
)

func main() {
	store := store.NewVectorStore()
	store.InsertVector([]float32{1, 2, 3})
	store.InsertVector([]float32{4, 5, 6})
	store.InsertVector([]float32{7, 8, 9})
	store.InsertVector([]float32{10, 11, 12})
	// vecs := store.GetAllVectors()
	// for _, v := range vecs {
	// 	fmt.Println(v)
	// }
	tar := []float32{1, 2, 3}
	neighbours := store.GetKNearestNeighbors(tar, 5)

	for _, v := range neighbours {
		fmt.Println(v)
	}
}
