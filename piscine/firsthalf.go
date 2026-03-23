package piscine

// import "github.com/01-edu/z01"

// func RetainFirstHalf(str string) string {
// 	runes := []rune(str) // Convert string to runes to handle multi-byte characters
// 	length := len(runes)

// 	if length == 0 {
// 		return "" // empty string case
// 	}

// 	if length == 1 {
// 		return str // single character case
// 	}

// 	half := length / 2 // automatically rounds down if odd
// 	return string(runes[:half])
// }

func RetainFirstHalf(str string) string {
if len(str) == 0 {
return ""
}

if len(str) == 1 {
return str
}

length := len([]rune(str))
string1 := []rune(str)
 half := length/2

 return  string(string1[:half])
}