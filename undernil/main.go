package main

import "fmt"

func main() {
	var i interface{}
	var p *int

	fmt.Println(i == nil) // true
	fmt.Println(p == nil) // true

	i = p
	fmt.Println(i == nil)
}
