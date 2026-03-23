package main

import (
	"fmt"
)

func LastWord(s string) string {
    lastWord := ""
    currentWordStart := -1

    for i, chr := range s {
        // 1. Found the start of a word
        if chr != ' ' && currentWordStart == -1 {
            currentWordStart = i
        }

        // 2. Found the end of a word (hit a space)
        if chr == ' ' && currentWordStart != -1 {
            lastWord = s[currentWordStart:i]
            currentWordStart = -1 // Reset the "bucket" to look for the next word
        }
    }

    // 3. Final Check: If the string didn't end with a space, 
    // the very last word is still in our "currentWordStart" bucket.
    if currentWordStart != -1  {
        lastWord = s[currentWordStart:]
    }

    if lastWord == "" {
        return "\n"
    }

    return lastWord + "\n"
}

func main() {
	// We use Print because the \n is already inside the return string
	fmt.Print(LastWord("this         ...       is sparta, then again, maybe    not "))
	fmt.Print(LastWord("shina "))
	fmt.Print(LastWord(""))
	fmt.Print(LastWord(" lorem, ipsum "))
	
}