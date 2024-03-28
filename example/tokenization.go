package main

import "strings"

func CreateVocabulary(data []string) (map[string]int, map[string]int) {
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
	return vocabulary, wordIndex
}

func VectorizeText(text string, vocabulary map[string]int, wordIndex map[string]int) []float32 {
	tokens := strings.Fields(strings.ToLower(text))
	vector := make([]float32, len(vocabulary))
	for _, token := range tokens {
		if idx, exists := wordIndex[token]; exists {
			vector[idx]++
		}
	}
	return vector
}
