package main

import (
	"fmt"
)

func ThirdTimeIsACharm(str string) string {
	// 1. Check if the string is too short (less than 3 characters)
	// This also handles the empty string case.
	if len(str) < 3 {
		return "\n"
	}

	result := ""

	// 2. Loop through the string
	// We start at index 2 (the 3rd character) and skip 3 every time
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}

	// 3. Add the required newline to the final collection
	return result + "\n"
}


func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}