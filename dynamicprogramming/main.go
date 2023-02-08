package main

import (
	"fmt"
	"log"
)

var results = make(map[int]int)

var str string = "str"
var post []Post = []Post{{name: "tony"}, {name: "may"}}

type Post struct {
	name string
	// age  int
}
func init() {
	log.SetPrefix("Dynamic Programming --")
	// concept of solving a problem with a stored version solution not reinventing the wheel
	//in this case lets have look at fibonacci
}
func main() {
	log.Println()
	fmt.Println("resultings to ", Fibonacci(10))
	fmt.Println("resultings to ", FibonacciMap(10))
	fmt.Println("resultings to ", results)
	fmt.Println("resultings to ", FibonacciLoop(10))
	fmt.Println("---------->>>>>>>>", str, post)
}


func Fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
func FibonacciMap(n int) int {
	if n <= 2 {
		return 1
	}
	if result, ok := results[n]; ok {
		return result
	}
	result := Fibonacci(n-1) + Fibonacci(n-2)
	results[n] = result
	return result
}
func FibonacciLoop(n int) []int {
	f := make([]int, n+1, n+2)
	if n < 2 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f
}
