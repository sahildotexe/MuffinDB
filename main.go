package main

import (
	"fmt"
	"math"
	"sort"
)

type Vector [3]float32

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

func (_ Leaf) isKDTreeNode() {}

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

func (kdtree KDTree) nearestNeighbor(query Vector) *Vector {
	_, nearest := kdtree.nearest(query, kdtree.Root, nil, math.MaxFloat32)
	return nearest
}

func (kdtree KDTree) nearest(query Vector, node KDTreeNode, best *Vector, bestDist float32) (float32, *Vector) {
	switch n := node.(type) {
	case Leaf:
		dist := euclideanDistance(query, n.Point)
		if dist < bestDist {
			return dist, &n.Point
		}
		return bestDist, best
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
		updatedBestDist, updatedBest := kdtree.nearest(query, nextNode, best, bestDist)
		if math.Abs(float64(query[n.SplitDimension]-n.SplitValue)) < float64(updatedBestDist) {
			return kdtree.nearest(query, otherNode, updatedBest, updatedBestDist)
		}
		return updatedBestDist, updatedBest
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

func main() {
	points := []Vector{{1.0, 2.0, 3.0}, {4.0, 5.0, 6.0}, {2.0, 3.0, 4.0}} // Add more points here
	kdTree := KDTree{Root: build(points, 0)}

	if nearest := kdTree.nearestNeighbor(Vector{1.0, 5.0, 3.0}); nearest != nil {
		fmt.Println("Nearest neighbor:", *nearest)
	}
}
