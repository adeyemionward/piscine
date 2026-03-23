package main

import "fmt"

func main() {
    // We use integers 0-9 instead of characters '0'-'9'
    for a := 9; a >= 0; a-- {
        for b := a - 1; b >= 0; b-- {
            for c := b - 1; c >= 0; c-- {
                
                // Print the combination
                fmt.Print(a, b, c)

                // Check: Are we at the very last combination?
                // If NOT, print the comma and space
                if !(a == 2 && b == 1 && c == 0) {
                    fmt.Print(", ")
                }
            }
        }
    }
    // Print one final newline at the very end
    fmt.Println()
}