package main

import (
	"fmt"
	"math"
	"sort"
)

type Vector []float32

type KDTreeNode interface {
	isKDTreeNode()
}

type Leaf struct {
	Point Vector
}

type Internal struct {
	Left           KDTreeNode
	Right          KDTreeNode
	SplitValue     float32
	SplitDimension int
}

func (_ Leaf) isKDTreeNode()     {}
func (_ Internal) isKDTreeNode() {}

type KDTree struct {
	Root KDTreeNode
}

func build(points []Vector, depth int) KDTreeNode {
	if len(points) == 1 {
		return Leaf{points[0]}
	}
	dim := depth % len(points[0]) // Assuming all points have the same dimension
	sortedPoints := make([]Vector, len(points))
	copy(sortedPoints, points)
	sortByDimension(sortedPoints, dim)
	medianIdx := len(sortedPoints) / 2
	medianValue := sortedPoints[medianIdx][dim]
	return Internal{
		Left:           build(sortedPoints[:medianIdx], depth+1),
		Right:          build(sortedPoints[medianIdx:], depth+1),
		SplitValue:     medianValue,
		SplitDimension: dim,
	}
}

func sortByDimension(points []Vector, dim int) {
	sort.Slice(points, func(i, j int) bool {
		return points[i][dim] < points[j][dim]
	})
}

func (kdtree KDTree) nearestNeighbor(query Vector, k int) []HeapVector {
	if k <= 0 {
		return nil
	}
	neighbors, _ := kdtree.nearest(query, kdtree.Root, make([]HeapVector, 0, k), math.MaxFloat32, k)
	return neighbors
}

func (kdtree KDTree) nearest(query Vector, node KDTreeNode, neighbors []HeapVector, bestDist float32, k int) ([]HeapVector, float32) {
	switch n := node.(type) {
	case Leaf:
		dist := euclideanDistance(query, n.Point)
		neighbors = pushHeap(neighbors, n.Point, dist, k)
		if len(neighbors) == k {
			bestDist = neighbors[0].distance
		}
		return neighbors, bestDist
	case Internal:
		var nextNode KDTreeNode
		var otherNode KDTreeNode
		if query[n.SplitDimension] < n.SplitValue {
			nextNode = n.Left
			otherNode = n.Right
		} else {
			nextNode = n.Right
			otherNode = n.Left
		}
		neighbors, bestDist = kdtree.nearest(query, nextNode, neighbors, bestDist, k)
		if math.Abs(float64(query[n.SplitDimension]-n.SplitValue)) < float64(bestDist) {
			neighbors, bestDist = kdtree.nearest(query, otherNode, neighbors, bestDist, k)
		}
		return neighbors, bestDist
	default:
		panic("unexpected node type")
	}
}

func euclideanDistance(p1, p2 Vector) float32 {
	var sum float32
	for i := 0; i < len(p1); i++ {
		sum += (p1[i] - p2[i]) * (p1[i] - p2[i])
	}
	return float32(math.Sqrt(float64(sum)))
}

type HeapVector struct {
	Point    Vector
	distance float32
}

type HeapVectors []HeapVector

func (h HeapVectors) Len() int           { return len(h) }
func (h HeapVectors) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h HeapVectors) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func pushHeap(neighbors []HeapVector, point Vector, dist float32, k int) []HeapVector {
	heap := make(HeapVectors, len(neighbors), k)
	copy(heap, neighbors)
	heap = heap[:min(len(heap)+1, k)]
	heap[len(heap)-1] = HeapVector{Point: point, distance: dist}
	sort.Sort(heap) // Using sort.Sort instead of heap.Sort()
	return heap
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	points := []Vector{{1.0, 2.0, 3.0, 5.0}, {4.0, 5.0, 6.0, 2.0}, {2.0, 3.0, 4.0, 1.5}, {1.0, 5.0, 3.0, 2.3}}
	kdTree := KDTree{Root: build(points, 0)}
	neighbors := kdTree.nearestNeighbor(Vector{2, 3, 4, 1.5}, len(points))
	k := 1
	for i := 0; i < k; i++ {
		fmt.Println("Point:", neighbors[i].Point, "Distance:", neighbors[i].distance)
	}
}
