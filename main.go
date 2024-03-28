package main

import (
	"GoVectorDB/kdtree"
	"fmt"
)

func main() {
	tree := kdtree.NewKDTree()
	tree.Insert(kdtree.Vector{ID: 1, Values: []float32{8, 2, 3, 4}})
	tree.Insert(kdtree.Vector{ID: 2, Values: []float32{2, 3, 4, 1.5}})
	tree.Insert(kdtree.Vector{ID: 3, Values: []float32{3, 4, 1.5, 2}})
	tree.Insert(kdtree.Vector{ID: 4, Values: []float32{4, 1.5, 2, 3}})
	tree.Insert(kdtree.Vector{ID: 5, Values: []float32{3, 2, 3, 4}})
	// tree.PrintTree()
	target := kdtree.Vector{Values: []float32{2, 3, 4, 1.5}}
	neighbors := tree.GetNeighbours(target, 5)
	k := 2
	for i := 0; i < k; i++ {
		fmt.Printf("ID: %d, Point: %v, Distance: %.02f\n", neighbors[i].Point.ID, neighbors[i].Point.Values, neighbors[i].Distance)
	}
}
