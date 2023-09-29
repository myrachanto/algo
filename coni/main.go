package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for range list {
		fmt.Println("--------------")
	}
}
