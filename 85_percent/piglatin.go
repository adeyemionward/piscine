package main

import (
	"os"
	"fmt"
)

func isVowel(r byte) bool {
	vowels := "aeiouAEIOU"
	for i := 0; i < len(vowels); i++ {
		if r == vowels[i] {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]

	// Rule: If number of arguments is not exactly one, print nothing
	if len(args) != 1 {
		return
	}

	word := args[0]
	firstVowelIndex := -1

	// Find the index of the first vowel
	for i := 0; i < len(word); i++ {
		if isVowel(word[i]) {
			firstVowelIndex = i
			break
		}
	}

	// Case 1: No vowels in the word
	if firstVowelIndex == -1 {
		fmt.Println("No vowels")
		return
	}

	// Case 2: Starts with a vowel (index 0)
	if firstVowelIndex == 0 {
		fmt.Printf("%say\n", word)
		return
	}

	// Case 3: Starts with consonants
	// Split the word: [vowel to end] + [start to vowel] + "ay"
	prefix := word[:firstVowelIndex]
	suffix := word[firstVowelIndex:]
	fmt.Printf("%s%say\n", suffix, prefix)
}