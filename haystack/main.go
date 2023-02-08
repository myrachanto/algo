package main

import "fmt"

func BinarySearch(haystack []int, needle int) int {
	low := 0
	high := len(haystack) - 1
	for low <= high {
		mid := (low + high) / 2
		if haystack[mid] < needle {
			low = mid + 1
		} else if haystack[mid] > needle {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
func generate() (res []int) {
	for i := 0; i < 200000; i++ {
		if i%2 == 0 {
			res = append(res, i)
		}
	}
	return
}

func main() {
	res := BinarySearch(generate(), 60002)
	fmt.Println("the results ", res)
}
