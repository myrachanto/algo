package main

import (
	"fmt"
	"sort"
)

func main() {
	var list = []int{2, 34, 56, 76, 78, 98, 87, 23, 54, 45, 76, 72, 20}
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	results := algosearch(98, list)
	fmt.Println(results)
}

func algosearch(needle int, hay []int) bool {
	low := 0
	high := len(hay) - 1
	for low <= high {
		median := (low + high)/2
		if hay[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(hay) || hay[low] != needle {
		return false
	}
	return true
}
