package main

import (
	"fmt"
)

func reverseNumber(n int) int {
	reversed := 0
	for n != 0 {
		digit := n % 10                // Extract the last digit
		reversed = reversed*10 + digit // Append the digit to the reversed number
		n /= 10                        // Remove the last digit from the number
	}
	return reversed
}

func main() {
	num := 12345
	reversed := reverseNumber(num)
	fmt.Println("Reversed:", reversed) // Output: Reversed: 54321
}
