package main
import (
	"fmt"	
)

func HashCode(dec string) string {	
	length := len(dec)
	hashed := ""
	for _, chr := range dec {
		h := int(chr) + length % 127
		if (chr < 32){
			h+= 33
		}
		hashed += string(rune(h))
	}
	return hashed	
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}