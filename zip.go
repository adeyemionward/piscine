package main

import "fmt"

func ZipString(s string) string {
	if s == "" {
		return ""
	}

	result := ""
	count := 1

	// We loop from the first character to the second-to-last
	for i := 0; i < len(s)-1; i++ {
		// Compare current character with the next one
		if s[i] == s[i+1] {
			count++
		} else {
			// They are different! Save the count and the character
			result += fmt.Sprintf("%d%c", count, s[i])
			// Reset for the new character
			count = 1
		}
	}

	// Step 3: Add the very last character group
	result += fmt.Sprintf("%d%c", count, s[len(s)-1])

	return result
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}