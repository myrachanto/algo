package main

import "fmt"

func main() {
	a := 6
	defer fmt.Println(&a)
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	c := 7

	a = c
	fmt.Println(&a, &c)
}
