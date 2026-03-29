package main

import (
	"fmt"
)

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	// 1. Remove ALL original spaces first
	clean := ""
	for _, c := range str {
		if c != ' ' {
			clean += string(c)
		}
	}

	// 2. Check the length of the CLEAN string
	if len(clean) < 5 {
		return "Invalid Input\n"
	}

	result := ""
	count := 0
	// 3. Use i loop to allow manual skipping
	for i := 0; i < len(clean); i++ {
		result += string(clean[i])
		count++

		if count == 5 {
			// If there's a character at i+1, it's the 6th one. Skip it!
			if i+1 < len(clean) {
				result += " "
				i++ // The Skip
			}
			count = 0
		}
	}

	return result + "\n"
}

func main() {
	fmt.Print(FifthAndSkip(" abcdefghijklmnopqrstuwxyz "))
	fmt.Print(FifthAndSkip("This is a short sentence "))
	fmt.Print(FifthAndSkip("abcdefg"))

	fmt.Print(FifthAndSkip("a b c f g h i j k "))
}
