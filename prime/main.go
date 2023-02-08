package main

import (
	"fmt"
	"math"
)

func Primer(num1, num2 int) []int {
	primes := []int{}
	if num1 < 2 || num2 < 2 {
		fmt.Println("num1 and num2 should be greator than 2")
	}
	for num1 <= num2 {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(num1))); i++ {
			if num1%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, num1)
		}
		num1++
	}
	return primes
}
func main() {
	fmt.Println(Primer(2, 100))
}
