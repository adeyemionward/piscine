package main

import (
	// "fmt"
)

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	sign := 1
	start := 0

	// Check for a leading sign
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	// Iterate through the remaining characters
	for i := start; i < len(s); i++ {
		// If any character is not a digit, return 0
		if s[i] < '0' || s[i] > '9' {
			return 0
		}

		// Convert character to int and add to result
		digit := int(s[i] - '0')
		result = result*10 + digit
	}

	return result * sign

}

// func main() {
	
	
//     fmt.Println(Atoi("12345"))
// 	fmt.Println(Atoi("0000000012345"))
// 	fmt.Println(Atoi("012 345"))
// 	fmt.Println(Atoi("Hello World!"))
// 	fmt.Println(Atoi("+1234"))
// 	fmt.Println(Atoi("-1234"))
// 	fmt.Println(Atoi("++1234"))
// 	fmt.Println(Atoi("--1234"))
// }


