package main

import (
	"os"
	"fmt"
)

func main() {
	// Step 1: Check argument count
	// os.Args[0] is the program name; we need exactly one more argument.
	if len(os.Args) != 2 {
		return
	}

	str := os.Args[1]
	if str == "" {
		return
	}

	// Step 2: Manually split the string into words
	// Since the subject says words are separated by exactly one space
	var words []string
	currentWord := ""

	for _, char := range str {
		if char == ' ' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}
	// Append the last word since there is no space at the end
	if currentWord != "" {
		words = append(words, currentWord)
	}

	// Step 3: Print words in reverse order
    for i := len(words) - 1; i >= 0; i-- {
        // Print the word directly (fmt knows how to print strings)
        fmt.Print(words[i])

        // Print a space between words, using double quotes " " for a string
        if i > 0 {
            fmt.Print(" ")
        }
    }

    // Step 4: Final newline using double quotes
    fmt.Print("\n")

	// Step 4: Final newline as required
	// fmt.Println()
}