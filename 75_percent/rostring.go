package main

import (
	"os"
	"fmt"
)

func main() {
	// Step 1: Check if there is exactly one argument
	if len(os.Args) != 2 {
		if len(os.Args) != 1 { // Only print newline if it's not just the program name
			fmt.Println()
		}
		return
	}

	str := os.Args[1]
	var words []string
	currentWord := ""

	// Step 2: Extract words (Handling multiple spaces)
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
	// Don't forget the last word!
	if currentWord != "" {
		words = append(words, currentWord)
	}

	// Step 3: Handle the rotation and printing
	if len(words) > 0 {
		// Print from the second word (index 1) to the end
		for i := 1; i < len(words); i++ {
			for _, r := range words[i] {
				fmt.Print(string(r))
			}
			// fmt.Print(words[i]) // fmt.Print can print the whole word at once
          
			fmt.Print(" ") // Space after each middle word
		}

		// Finally, print the first word (index 0)
		for _, r := range words[0] {
			fmt.Print(string(r))
		}
	}

	// Step 4: Always end with a newline
	fmt.Print("\n")
}