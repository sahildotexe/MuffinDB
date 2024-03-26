package main

import (
	"fmt"
	"math"
)

type Vector [3]float64

type VectorDB struct {
	vectors []Vector
}

func NewVectorDB() *VectorDB {
	return &VectorDB{}
}

func (db *VectorDB) AddVector(v Vector) {
	db.vectors = append(db.vectors, v)
}

func (db *VectorDB) GetVector(i int) Vector {
	return db.vectors[i]
}

func (db *VectorDB) FindClosest(v Vector) Vector {
	if len(db.vectors) == 0 {
		return Vector{}
	}

	closest := db.vectors[0]
	closestDist := euclideanDistance(v, closest)
	for _, vector := range db.vectors {
		dist := euclideanDistance(v, vector)
		if dist < closestDist {
			closest = vector
			closestDist = dist
		}
	}
	return closest
}

func euclideanDistance(v1, v2 Vector) float64 {
	return math.Sqrt((v1[0]-v2[0])*(v1[0]-v2[0]) + (v1[1]-v2[1])*(v1[1]-v2[1]) + (v1[2]-v2[2])*(v1[2]-v2[2]))
}

func main() {
	db := NewVectorDB()
	db.AddVector(Vector{1, 2, 3})
	db.AddVector(Vector{4, 5, 6})

	// get vector by index.
	v2 := db.GetVector(1)
	fmt.Println(v2)

	v := Vector{2.0, 5.0, 6.0}
	closest := db.FindClosest(v)
	fmt.Println(closest)
}
