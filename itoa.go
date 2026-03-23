
package main
import (
	"fmt"
)

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	s := ""
	for n > 0 {
		digit := n % 10
		s = string(rune('0'+digit)) + s // prepend digit
		n = n / 10
	}

	if negative {
		s = "-" + s
	}

	return s
}

func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}