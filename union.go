package main

import (
	"os"
	"github.com/01-edu/z01"
)

func alreadySeen(str string, c rune, index int) bool {
	for i := 0; i < index; i++ {
		if rune(str[i]) == c {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	// check first string
	for i, c := range s1 {
		if !alreadySeen(s1, c, i) {
			z01.PrintRune(c)
		}
	}

	// check second string
	for i, c := range s2 {
		if !alreadySeen(s1+s2, c, len(s1)+i) {
			z01.PrintRune(c)
		}
	}

	z01.PrintRune('\n')
}