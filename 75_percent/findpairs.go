
// FIND PAIRS
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 1. Basic Argument Check
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	argArr := os.Args[1]
	argTarget := os.Args[2]

	// 2. Manual Bracket Check
	if len(argArr) < 2 || argArr[0] != '[' || argArr[len(argArr)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}

	// 3. The Manual Scanner
	var nums []int
	currentNumStr := ""

	// Loop through everything INSIDE the brackets
	for i := 1; i < len(argArr)-1; i++ {
		char := argArr[i]

		// If it's a number or a minus sign, add it to our temporary string
		if (char >= '0' && char <= '9') || char == '-' {
			currentNumStr += string(char)
		} else if char == ',' {
			// We hit a comma! Time to convert the number we found
			if currentNumStr != "" {
				n, err := strconv.Atoi(currentNumStr)
				if err != nil {
					fmt.Printf("Invalid number: %s\n", currentNumStr)
					return
				}
				nums = append(nums, n)
				currentNumStr = "" // Empty the "holding area"
			}
		} else if char == ' ' {
			continue // Just ignore spaces
		} else {
			// If we see 'p' or anything else that isn't a number/comma/space
			fmt.Printf("Invalid number: %c\n", char)
			return
		}
	}

	// 4. Grab the LAST number (because there is no comma after the last number!)
	if currentNumStr != "" {
		n, err := strconv.Atoi(currentNumStr)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", currentNumStr)
			return
		}
		nums = append(nums, n)
	}

	// 5. Convert the Target Sum
	targetSum, err := strconv.Atoi(argTarget)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}

	// 6. The "Two-Finger Search" (Logic stays the same!)
	var pairs [][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == targetSum {
				pairs = append(pairs, []int{i, j})
			}
		}
	}

	// 7. Final Result
	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
	} else {
		fmt.Printf("Pairs with sum %d: %v\n", targetSum, pairs)
	}
}