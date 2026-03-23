package main

import "fmt"

// FindPrevPrime returns the largest prime number <= nb, or 0 if none exists
func FindPrevPrime(nb int) int {
	// Check each number from nb down to 2
	for i := nb; i >= 2; i-- {
		isPrime := true
		// Check if i is divisible by any number from 2 to i-1
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(FindPrevPrime(5))  // 5
	fmt.Println(FindPrevPrime(4))  // 3
	fmt.Println(FindPrevPrime(1))  // 0
	fmt.Println(FindPrevPrime(10)) // 7
}