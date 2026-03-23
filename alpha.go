package main

import (
	"fmt"
)



func RepeatAlpha(s string) string {
	result := ""
	

	for _, c:= range s {
		repeat := 1
	    if(c >= 'a' && c <= 'z'){
			repeat = int(c-'a') + 1
	    }

	    if(c >= 'A' && c <= 'Z'){
			repeat = int(c-'A') + 1
	    }
	     
        for j := 1;  j <= repeat; j++ {
			result += string(c)
	    }
	}
	return result
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha("!!!"))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}