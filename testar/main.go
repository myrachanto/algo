package main

import (
	"fmt"
	"reflect"
	"testing"
)

func assert(t *testing.T, a, b any) {
	if reflect.DeepEqual(a, b) {
		t.Errorf("%v != %v", a, b)
	}
}
func deleteItemFromSlice(index int, arr []int) []int {
	arr[len(arr)-1] = arr[index]
	arr = arr[:len(arr)-1]
	return arr
}

func deleteItemFromSlice2(index int, arr []int) []int {
	var res []int
	for k, v := range arr {
		if k != index {
			res = append(res, v)
		}
	}
	return res
}
func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr1, len(arr1))
	arr2 := deleteItemFromSlice(4, arr1)
	fmt.Println(arr2, len(arr2))
	arr3 := deleteItemFromSlice2(4, arr1)
	fmt.Println(arr3, len(arr3))
}
