package main

import (
	"fmt"
)


func Slice(a []string, nbrs ...int) []string {
	l := len(a)
	if l == 0 {
		return nil
	}

	start := nbrs[0]
	end := l

	// If they gave us a second number, use it as the end
	if len(nbrs) > 1 {
		end = nbrs[1]
	}

	// Turn negative "secret codes" into real positions
	if start < 0 {
		start = l + start
	}
	if end < 0 {
		end = l + end
	}

	// Safety checks: don't go out of bounds!
	if start < 0 {
		start = 0
	}
	if end > l {
		end = l
	}

	// If start is after end, it's impossible to cut!
	if start >= end {
		return nil
	}

	return a[start:end]
}
func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n",Slice(a, 1))
	fmt.Printf("%#v\n",Slice(a, 2, 4))
	fmt.Printf("%#v\n",Slice(a, -3))
	fmt.Printf("%#v\n",Slice(a, -1, -2))
	fmt.Printf("%#v\n",Slice(a, 2, 0))
}
