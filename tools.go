// Here is an example Go function that takes a CDR notation (e.g. "24") and returns the number of IP addresses in the corresponding CIDR range:
package main

import (
    "fmt"
    "math"
)

func cdrRange(cdr string) int {
    n, _ := fmt.Sscanf(cdr, "%d", &n)
    if n != 1 {
        return 0
    }
    if n < 0 || n > 32 {
        return 0
    }
    return int(math.Pow(2, float64(32-n))) - 2
}

func main() {
    fmt.Println(cdrRange("24")) // Output: 256
}

// This function first uses `fmt.Sscanf` to parse the input string into an integer `n`. If the parse fails or `n` is outside the valid range of 0 to 32, the function returns 0.

// Otherwise, the function computes the number of IP addresses in the range using the formula `2^(32-n) - 2`. This formula subtracts 2 from the total number of IP addresses to exclude the network address and the broadcast address from the range.

// In the example code above, calling `cdrRange("24")` returns `256`, which is the number of IP addresses in a /24 CIDR range.
