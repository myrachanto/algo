package main

import "fmt"

func primer(limit int) []int {
	primes := []int{2}
	for i := 3; i <= limit; i++ {
		if !filter(i, primes) {
			primes = append(primes, i)
		}
	}
	return primes
}
func filter(n int, m []int) bool {
	var res bool
	for _, v := range m {
		if n%v == 0 {
			res = true
		}
	}
	return res
}
func main() {
	fmt.Println(primer(100))
}
