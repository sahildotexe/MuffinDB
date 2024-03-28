package main

import (
	"GoVectorDB/store"
	"fmt"
	"strings"
)

func main() {
	store := store.NewVectorStore()

	data := []string{
		"Cricket is a popular sport in India",
		"Virat Kohli represents India in international cricket",
		"Virat Kohli plays for RCB in IPL",
		"Virat Kohli is my favorite cricketer",
	}

	// Tokenization and Vocabulary Creation
	vocabulary := make(map[string]int)
	wordIndex := make(map[string]int)
	index := 0
	for _, sentence := range data {
		tokens := strings.Fields(strings.ToLower(sentence))
		for _, token := range tokens {
			if _, exists := vocabulary[token]; !exists {
				vocabulary[token] = index
				wordIndex[token] = index
				index++
			}
		}
	}

	// Vectorization
	sentenceVectors := make(map[string][]float32)
	for _, sentence := range data {
		tokens := strings.Fields(strings.ToLower(sentence))
		vector := make([]float32, len(vocabulary))
		for _, token := range tokens {
			vector[wordIndex[token]]++
		}
		sentenceVectors[sentence] = vector
	}

	// Inserting Vectors into the Store
	for sentence, vector := range sentenceVectors {
		store.InsertVector(sentence, vector)
	}

	// Searching for Similarity
	query := "Which team does Virat Kohli play for in IPL?"

	queryTokens := strings.Fields(strings.ToLower(query))
	queryVector := make([]float32, len(vocabulary))
	for _, token := range queryTokens {
		if idx, exists := wordIndex[token]; exists {
			queryVector[idx]++
		}
	}

	fmt.Println("Query Vector:", queryVector)

	vecs := store.GetAllVectors()
	for _, v := range vecs {
		fmt.Println(v)
	}
	fmt.Println("-------------------")
	neighbours := store.GetKNearestNeighbors(queryVector, 4)

	for _, v := range neighbours {
		fmt.Printf("Text: %s, Distance: %f\n", v.Point.Text, v.Distance)
	}
}
