package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	arr := []int{1, 2, 4, 5, 3, 7, 8, 9, 6}
	sort.Ints(arr)
	fmt.Println(arr)
	Reversing(arr)
	fmt.Println(arr)
	Reversing2(arr)
	fmt.Println(arr)
	// shufle(arr)
	// fmt.Println(arr)
	batches := Batching(arr, 4)
	fmt.Println(batches)
	// sort.Ints()
	a := []int{3, 2, 1, 4, 3, 2, 1, 4, 1, 5, 5, 6, 7, 6} // any item can be sorted
	res := inpalaceDups(a)
	fmt.Println(res)
	as := []int{1, 2, 3, 4, 5}
	bs := []int{0, 0, 0}
	copy(bs, as)
	fmt.Println("----------------", bs)

}
func sorting(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
func Reversing(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}
func Reversing2(a []int) {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}
func shufle(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
func Batching(a []int, batchSize int) [][]int {
	batches := make([][]int, 0, (len(a)+batchSize-1)/batchSize)

	for batchSize < len(a) {
		a, batches = a[batchSize:], append(batches, a[0:batchSize:batchSize])
	}
	batches = append(batches, a)
	return batches
}
func inpalaceDups(a []int) []int {
	sort.Ints(a)
	j := 0
	for i := 1; i < len(a); i++ {
		if a[j] == a[i] {
			continue
		}
		j++
		a[j] = a[i]
	}
	results := a[:j+1]
	return results
}
