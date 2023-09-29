package main

import "testing"

// go test -bench=.
// go test -bench=. -benchtime 10s
func BenchmarkCompareIfItemExist(b *testing.B) {
	arr1 := []string{"a", "b", "c", "d", "e"}
	arr2 := []string{"m", "n", "o", "p"}
	for i := 0; i < b.N; i++ {
		CompareIfItemExist(arr1, arr2)
	}

}
func BenchmarkCompareIfItemExist2(b *testing.B) {
	arr1 := []string{"a", "b", "c", "d", "e"}
	arr2 := []string{"m", "n", "o", "p"}
	for i := 0; i < b.N; i++ {
		CompareIfItemExist2(arr1, arr2)
	}

}
