package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ls := []int{3, 2, 6, 4, 7, 1, 9, 8, 0, 5}
	results := BubbleSort(ls)
	fmt.Println(results)
}

func MergeSort(ls []int) []int {
	if len(ls) < 2 {
		return ls
	}
	mid := len(ls) / 2
	return Merge2(MergeSort(ls[:mid]), MergeSort(ls[mid:]))
}
func Merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	ls := make([]int, size)
	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			ls[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			ls[k] = left[i]
			i++
		} else if left[i] < right[j] {
			ls[k] = left[i]
			i++
		} else {
			ls[k] = right[j]
			j++
		}
	}
	return ls
}
func Merge2(left, right []int) []int {
	ls, i, j := []int{}, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			ls = append(ls, left[i])
			i++
		} else {
			ls = append(ls, right[j])
			j++
		}
	}
	for ; i < len(left); i++ {
		ls = append(ls, left[i])
	}
	for ; j < len(right); j++ {
		ls = append(ls, right[j])
	}
	return ls

}
func BubbleSort(ls []int) []int{
	n := len(ls)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++{
			if ls[i]> ls[i+1]{
				ls[i], ls[i+1] = ls[i+1],ls[i]
				swapped = true
			}
		}
	}
	return ls
}
func QuickSort(ls []int) []int {
	if len(ls) < 2 {
		return ls
	}
	left, right := 0, len(ls)-1
	pivot := rand.Int() % len(ls)
	ls[pivot], ls[right] = ls[right], ls[pivot]
	for i, _ := range ls{
		if ls[i] < ls[right]{
			ls[left],ls[i] = ls[i],ls[left]
			left++
		}
	}
	ls[left],ls[right] = ls[right],ls[left]
	QuickSort(ls[:left])
	QuickSort(ls[left+1:])
	return ls
}


// func trymergerSort(listA, listB []int) []int {
// 	var results []int
// 	for i := 0; i < len(listA); i++ {
// 		for j := 0; j < len(listB); j++ {
// 			if i < j {
// 				results = append(results, i)
// 			} else {
// 				results = append(results, j)
// 			}
// 		}
// 	}
// 	return results
// }
