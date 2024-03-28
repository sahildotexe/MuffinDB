package main

import (
	"GoVectorDB/store"
	"fmt"
	"strings"
)

func main() {
	store := store.NewVectorStore()

	sentences := []string{
		"I eat mango",
		"mango is my favorite fruit",
		"mango, apple, oranges are fruits",
		"fruits are good for health",
	}

	// Tokenization and Vocabulary Creation
	vocabulary := make(map[string]int)
	for _, sentence := range sentences {
		tokens := strings.Fields(strings.ToLower(sentence))
		for _, token := range tokens {
			vocabulary[token]++
		}
	}

	// Assign unique indices to words in the vocabulary
	wordToIndex := make(map[string]int)
	index := 0
	for word := range vocabulary {
		wordToIndex[word] = index
		index++
	}

	// Vectorization
	var sentenceVectors map[string][]float32 = make(map[string][]float32)
	for _, sentence := range sentences {
		tokens := strings.Fields(strings.ToLower(sentence))
		vector := make([]float32, len(vocabulary))
		for _, token := range tokens {
			vector[wordToIndex[token]]++
		}
		sentenceVectors[sentence] = vector
	}

	for sentence, vector := range sentenceVectors {
		store.InsertVector(sentence, vector)
	}

	// Searching for Similarity
	querySentence := "Mango is the best fruit"
	queryVector := make([]float32, len(vocabulary))
	queryTokens := strings.Fields(strings.ToLower(querySentence))
	for _, token := range queryTokens {
		if index, ok := wordToIndex[token]; ok {
			queryVector[index]++
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
