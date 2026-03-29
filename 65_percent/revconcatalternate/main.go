package main

import (
	"fmt"
)


func RevConcatAlternate(slice1, slice2 []int) []int {
	res := []int{}
	len1 := len(slice1)
	len2 := len(slice2)

	// Keep going as long as there is a block in either stack
	for len1 > 0 || len2 > 0 {
		if len1 > len2 {
			// Slice1 is longer, take from the end of Slice1
			res = append(res, slice1[len1-1])
			len1--
		} else if len2 > len1 {
			// Slice2 is longer, take from the end of Slice2
			res = append(res, slice2[len2-1])
			len2--
		} else {
			// They are equal! Take from Slice1 then Slice2
			res = append(res, slice1[len1-1])
			res = append(res, slice2[len2-1])
			len1--
			len2--
		}
	}
	return res
}
func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
