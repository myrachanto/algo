package main

import "fmt"

type res struct{ k, v int }

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
		if v*v > n {
			break
		}
		if n%v == 0 {
			res = true
		}
	}
	return res
}
func findSumInts(target int, arr []int) []res {
	ress := []res{}
	found := make(map[int]bool)
	for _, v := range arr {
		complement := target - v
		if found[complement] {
			ress = append(ress, res{v, complement})
		}
		found[v] = true
	}
	return ress

}

func main() {
	target := 18
	primes := primer(target)
	res := findSumInts(target, primes)
	fmt.Println(res)
}
