package main

import (
	"fmt"
	"strings"
	"github.com/sahildotexe/swah-db/swah"
)

func main() {
	// Connect to the Vector Store
	store := swah.Connect()

	// Sample Data to Insert
	data := []string{
		"Cricket is a popular sport in India",
		"Virat Kohli represents India in international cricket",
		"Virat Kohli plays for RCB in IPL",
		"Virat Kohli is my favorite cricketer",
	}

	// Create Vocabulary and Word Index
	vocabulary, wordIndex := CreateVocabulary(data)

	// Vectorization
	vectors := make(map[string][]float32)
	for _, sentence := range data {
		vector := VectorizeText(sentence, vocabulary, wordIndex)
		vectors[sentence] = vector
	}

	// Inserting Vectors into the Store
	for sentence, vector := range vectors {
		store.InsertVector(sentence, vector)
	}

	// Get top 3 similar sentences
	query := "Which team does Virat Kohli play for in IPL?"
	queryVector := VectorizeText(strings.ToLower(query), vocabulary, wordIndex)
	fmt.Println("Query Prompt: ", query)
	k := 3
	neighbours := store.GetKNearestNeighbors(queryVector, k)
	fmt.Printf("\nTop %d Similar Sentences:\n", k)
	for _, v := range neighbours {
		fmt.Printf("Text: %s, Simlarity= %f\n", v.Point.Text, v.Distance)
	}
}
