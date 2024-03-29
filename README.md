# <div align="center"> <img src="https://github.com/sahildotexe/MuffinDB/blob/main/logo.png" alt="Image Alt Text" width="200"/>  </div>

MuffinDB is a simple and efficient vector store database written in Go. It uses a KD-Tree data structure to store and retrieve vectors efficiently, supporting basic operations like inserting, deleting, updating, and querying vectors based on proximity. With its optimized KD-Tree implementation, MuffinDB enables fast and accurate nearest neighbor searches, making it an ideal choice for efficiently storing and retrieving vectors in AI applications that require similarity-based queries or nearest neighbor lookups.


## Use Cases

MuffinDB, a vector store database, can be used in various applications and scenarios where efficient storage and retrieval of high-dimensional data is required. Some potential use cases include:

- **Similarity Search**
  - Content-based recommendation systems (e.g., movies, music, products)
  - Image and video similarity search
  - Document and text similarity search
  - Nearest neighbor search in high-dimensional spaces

- **Natural Language Processing (NLP)**
  - Storing and searching word embeddings
  - Semantic similarity between documents or sentences
  - Language modeling and text generation

- **Computer Vision**
  - Face recognition and facial feature matching
  - Image and object recognition and retrieval
  - Clustering and categorization of images

- **Bioinformatics**
  - Storing and analyzing gene sequences
  - Protein structure similarity search
  - Clustering and classification of biological data

- **Recommender Systems**
  - Building personalized recommendation engines
  - Collaborative filtering based on user/item embeddings
  - Content-based filtering using vector representations

- **Anomaly Detection**
  - Detecting anomalies or outliers in high-dimensional data
  - Fraud detection in financial transactions
  - Network intrusion detection

- **Data Clustering and Dimensionality Reduction**
  - Efficiently clustering high-dimensional data points
  - Visualizing and exploring high-dimensional data

These are just a few examples, and the applications of vector databases can extend to various domains where efficient storage, retrieval, and similarity search of high-dimensional data are required.

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
You can find the entire code at [Example](https://github.com/sahildotexe/MuffinDB/tree/main/example) . In this example, I've used Count Vectorization approach for vectorization but you can use any approach as per your choice.

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

