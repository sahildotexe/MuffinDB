package main

import (
	"GoVectorDB/kdtree"
	"fmt"
)

func main() {
	points := []kdtree.Vector{
		{ID: 1, Values: []float32{1.0, 2.0, 3.0, 5.0}},
		{ID: 2, Values: []float32{4.0, 5.0, 6.0, 2.0}},
		{ID: 3, Values: []float32{2.0, 3.0, 4.0, 1.5}},
		{ID: 4, Values: []float32{1.0, 5.0, 3.0, 2.3}},
	}
	tree := kdtree.KDTree{Root: kdtree.BuildTree(points, 0)}
	target := kdtree.Vector{Values: []float32{2, 3, 4, 1.5}}
	neighbors := tree.GetNeighbours(target, 4)
	k := 4
	for i := 0; i < k; i++ {
		fmt.Printf("ID: %d, Point: %v, Distance: %.02f\n", neighbors[i].Point.ID, neighbors[i].Point.Values, neighbors[i].Distance)
	}
}
