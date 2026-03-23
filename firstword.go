package main
import (
	"fmt"
	
)

func FirstWord(s string) string {
    start :=  -1
    for i, chr := range s {
		if chr != ' ' && start == -1 {
			start = i
		}

	
		if chr == ' ' && start != -1 {
			return s[start:i] + "\n"
		}
		
    }
	if start == -1 {
		return "\n"
	}


    return  s[start:] + "\n"
}


func main() {
	fmt.Print(FirstWord("hello there"))
	// fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}