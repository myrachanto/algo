package main

import (
	"fmt"
)

// 1. what does the following function prints?
func deferer() {
	for i := 0; i < 5; i++ {
		defer println(i)
	}
}

// 2. what is wrong with this function and how can you solve it?
func copier() {
	data := []int{1, 2, 3, 4, 5, 6}
	mapa := make([]int, 5)
	for _, d := range data {
		mapa = append(mapa, d)
	}
	fmt.Println(mapa)
}

// 3. what is the result of the following question?
func adder() {
	var num int8 = 127
	res := num + 1
	fmt.Println(res)
}

// 4. What will the function below print?
func slicer() {
	data := []int{1, 2, 3, 4, 5, 6}
	res := data[:3]
	final := res[2:4]
	fmt.Println(final)
}

// 5. what is wrong with the function below
func sliceCopier() {
	data := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	// slog.info()
	results := [][]byte{}
	for _, item := range data {
		res := make([]byte, len(item))
		// fmt.Println("---------", item)
		copy(make([]byte, len(item)), item[:])
		results = append(results, res)
	}
	fmt.Println(results)
}
func main() {
	deferer()
	copier()
	adder()
	slicer()
	sliceCopier()
}
