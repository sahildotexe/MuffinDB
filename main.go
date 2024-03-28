package main

import (
	"github.com/sahildotexe/swah-db/swah"
	"github.com/sahildotexe/swah-db/utils"
	"fmt"
	"strings"
)

func main() {
	store := swah.NewVectorStore()

	data := []string{
		"Cricket is a popular sport in India",
		"Virat Kohli represents India in international cricket",
		"Virat Kohli plays for RCB in IPL",
		"Virat Kohli is my favorite cricketer",
	}

	// Create Vocabulary and Word Index
	vocabulary, wordIndex := utils.CreateVocabulary(data)

	// Vectorization
	vectors := make(map[string][]float32)
	for _, sentence := range data {
		vector := utils.VectorizeText(sentence, vocabulary, wordIndex)
		vectors[sentence] = vector
	}

	// Inserting Vectors into the Store
	for sentence, vector := range vectors {
		store.InsertVector(sentence, vector)
	}

	// Searching for Similarity
	query := "Which team does Virat Kohli play for in IPL?"
	queryVector := utils.VectorizeText(strings.ToLower(query), vocabulary, wordIndex)
	fmt.Println("Query Prompt: ", query)
	// Get top 3 similar sentences
	k := 3
	neighbours := store.GetKNearestNeighbors(queryVector, k)
	fmt.Println("\nTop 3 Similar Sentences: ")
	for _, v := range neighbours {
		fmt.Printf("Text: %s, Simlarity= %f\n", v.Point.Text, v.Distance)
	}

}
