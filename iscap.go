package main

import (
	"fmt"
)

func IsCapitalized(s string) bool {
	if(len(s) == 0){
		return false
	}
	isWord := false
	for _, c := range s {

		if c != ' ' && !isWord {
				isWord = true
			if(c >= 'a' && c <= 'z') {
			return false
			}
		}
		if isWord && c == ' ' {
			isWord = false
		}

	}
	
	return true
}

func main() {
fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
