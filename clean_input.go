package main

import "strings"

func cleanInput(input string) []string {

	words := strings.Fields(input)
	cleanedWords := []string{}
	for _, word := range words {
		word := strings.ToLower(word)
		cleanedWords = append(cleanedWords, word)
	}
	return cleanedWords
}
