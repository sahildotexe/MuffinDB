# MuffinDB

MuffinDB is a simple and efficient vector store database written in Go. It uses a KD-Tree data structure to store and retrieve vectors efficiently. The database supports basic operations like inserting, deleting, updating, and querying vectors based on their proximity to a given query vector.

## Features

- **Fast Vector Insertion**: MuffinDB allows you to insert vectors quickly using the `InsertVector` method.
- **Efficient Nearest Neighbor Search**: The KD-Tree data structure enables fast nearest neighbor searches, allowing you to retrieve the `k` nearest vectors to a given query vector using the `GetKNearestNeighbors` method.
- **Vector Retrieval**: You can retrieve a specific vector by its ID using the `GetVector` method.
- **Vector Deletion**: MuffinDB supports deleting vectors from the database using the `DeleteVector` method.
- **Vector Updating**: You can update an existing vector in the database using the `UpdateVector` method.
- **Persistence**: MuffinDB stores the vector data on disk using Go's built-in `encoding/gob` package, ensuring data persistence across program restarts.

## Getting Started

To get started with MuffinDB, follow these steps:

1. Install Go on your machine if you haven't already.

2. Create a directory for your project
```bash
mkdir my-project
```

3. Navigate to the project directory:
```bash
cd my-project
```

4. Initialize a new Go module
```bash
go mod init example.com/my-project
```
This will create a go.mod file in your project directory.

5. Import MuffinDB package
```bash
go get github.com/sahildotexe/MuffinDB
```

## Usage

```go
package main

import (
	"fmt"
	"strings"

	"github.com/sahildotexe/MuffinDB/muffin"
)

func main() {
	// Connect to the Vector Store
	store := muffin.Connect()

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
```

Contributions to MuffinDB are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

