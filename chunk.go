package main

import (
	"fmt"
)

// Chunk splits a slice into sub-slices of a specific size.
func Chunk(slice []int, size int) {
	// 1. The "Zero" Guard
	if size <= 0 {
		fmt.Println()
		return
	}

	// 2. The "Master Container"
	var result [][]int

	// 3. The "Jumping" Loop
	for i := 0; i < len(slice); i += size {
		end := i + size

		// 4. The "Safety Cap"
		if end > len(slice) {
			end = len(slice)
		}

		// 5. The "Slicing" Action
		result = append(result, slice[i:end])
	}

	// 6. Final Output
	fmt.Println(result)
}

func main() {
	// Testing with different batch sizes
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
}