package main
import "fmt"


func WordFlip(str string) string {
	// Step 1: Manual split while ignoring extra spaces
	var words []string
	currentWord := ""

	for _, char := range str {
		if char == ' ' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}
	// Add the final word if the bucket isn't empty
	if currentWord != "" {
		words = append(words, currentWord)
	}

	// Step 2: Check if we actually found any words
	if len(words) == 0 {
		return "Invalid Output\n"
	}

	// Step 3: Build the reversed string
	result := ""
	for i := len(words) - 1; i >= 0; i-- {
		result += words[i]
		
		// Add a space between words, but not after the very last one
		if i > 0 {
			result += " "
		}
	}

	// Step 4: Return the final string with a newline
	return result + "\n"
}


func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}