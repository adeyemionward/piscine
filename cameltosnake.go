package main

import (
	"fmt"
)


func CamelToSnakeCase(s string) string{
	if(len(s) == 0){
		return ""
	}
	runes := []rune(s)
	result := ""
	
	for i, r := range s{
	 if ( r < 'a' || r > 'z') && (r < 'A' || r > 'Z' ) {
			return s
		}
	 if i > 0 && r >= 'A' && r <= 'Z' && s[i-1] >= 'A' && s[i-1] <= 'Z' {
			return s
		}
	}

	
	
	last := s[len(s)-1]
	if last >= 'A' && last <= 'Z' {
		return s
	}

	

	for i, r := range runes{
		if(r >= 'A' && r <= 'Z'){
			if(i != 0){
				result += "_"
			}
			result += string(r )
		}else{
			result += string(r)
		}		
	}

	return  result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCasE"))
	fmt.Println(CamelToSnakeCase("hey2"))
}