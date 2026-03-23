package main

import (
	"fmt"
)



func ConcatAlternate(slice1, slice2 []int) []int {
    // PHASE 1: The "Who Goes First?" Selection
    first := slice1
    second := slice2

    // Rule: Start with the longest. If equal, start with slice1.
    if len(slice2) > len(slice1) {
        first = slice2
        second = slice1
    }

    result := []int{}
    i := 0

    // PHASE 2: The Zipper (Alternating)
    // This loop runs ONLY as long as BOTH slices have items left.
    for i < len(first) && i < len(second) {
        result = append(result, first[i])  // Take from the Leader
        result = append(result, second[i]) // Take from the Follower
        i++ 
    }

    // PHASE 3: The Cleanup (Leftovers)
    // This loop picks up whatever is left in the longer slice.
    for i < len(first) {
        result = append(result, first[i])
        i++
    }

    return result
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10, 12}, []int{1, 3, 5, 7, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}