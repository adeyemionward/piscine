package main

import (
	"fmt"
	"strconv"
)

func NotDecimal(dec string) string {
    if dec == "" {
        return "\n"
    }

    dotIdx := -1
    hasDigits := false

    for i, num := range dec {
        if num == '.' {
            // Check if we already found a dot (only one allowed)
            if dotIdx != -1 {
                return dec + "\n"
            }
            dotIdx = i
            continue
        }
        if num == '-' {
            // Minus must be at the very start
            if i != 0 {
                return dec + "\n"
            }
            continue
        }
        if num < '0' || num > '9' {
            return dec + "\n" // Not a number
        }
        hasDigits = true
    }

    // If no digits or no dot, return as is
    if !hasDigits || dotIdx == -1 {
        return dec + "\n"
    }

    // Split it
    before := dec[:dotIdx]
    after := dec[dotIdx+1:]

    // SPECIAL RULE: If after the dot is ONLY zeros (like 174.0 or 174.000)
    onlyZeros := true
    for _, r := range after {
        if r != '0' {
            onlyZeros = false
            break
        }
    }
    
    // If it's just "174.0", return "174"
    if onlyZeros {
        return before + "\n"
    }

    // The "Glue": join them and convert to remove leading zeros (like "01" -> "1")
    // Keep the minus sign if it exists!
    combined := before + after
    
    // Atoi handles "-0125" -> "-125" automatically!
    val, _ := strconv.Atoi(combined)
    
    // BUT! strconv.Itoa(-0) becomes "0". 
    // If the original started with "-0.", Itoa might lose the "-"
    res := strconv.Itoa(val)
    if combined[0] == '-' && res[0] != '-' && val == 0 {
        return "-0\n" // Rare case edge case
    }

    return res + "\n"
}

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
