package piscine

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	hexChars := "0123456789abcdef"

	// --- HEX DISPLAY (4 per line) ---
	for i := 0; i < len(arr); i++ {

		firstDigit := arr[i] / 16
		secondDigit := arr[i] % 16

		z01.PrintRune(rune(hexChars[firstDigit]))
		z01.PrintRune(rune(hexChars[secondDigit]))

		// Space if not end of line or last element
		if (i+1)%4 != 0 && i != len(arr)-1 {
			z01.PrintRune(' ')
		}
        
		// New line every 4 bytes
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		}
	}

	// If last line not complete
	if len(arr)%4 != 0 {
		z01.PrintRune('\n')
	}

	// --- ASCII DISPLAY ---
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 32 && arr[i] <= 126 {
			z01.PrintRune(rune(arr[i]))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

