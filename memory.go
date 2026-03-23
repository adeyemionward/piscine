package piscine

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	hexChars := "0123456789abcdef"

	// --- STEP 1: HEXADECIMAL DISPLAY ---
	for i := 0; i < len(arr); i++ {
		// Calculate hex digits
		firstDigit := arr[i] / 16
		secondDigit := arr[i] % 16

		// Print the two characters from our hex "alphabet"
		z01.PrintRune(rune(hexChars[firstDigit]))
		z01.PrintRune(rune(hexChars[secondDigit]))

		// Add a space between hex values for readability
		// (Optional: check your specific example if the space is needed)
		if i < len(arr)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')

	// --- STEP 2: ASCII DISPLAY ---
	for i := 0; i < len(arr); i++ {
		// Rule: Only print if it's a "Graphic" (Printable) character
		// Space (32) up to Tilde (126)
		if arr[i] >= 32 && arr[i] <= 126 {
			z01.PrintRune(rune(arr[i]))
		} else {
			// Non-printable characters become a dot
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}