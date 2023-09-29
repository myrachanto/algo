package main

import "fmt"

func main() {
	// interview question
	// given 2 slices evaluate if an Item exist in both slices
	// if it exists return true else return false
	// example ["a","b","c", "d","e"], ["m", "n", "o", "p"] returns false
	// example ["a","b","c", "d","e"], ["m", "n", "o", "p", "e"] returns true
	testsCases := []struct {
		name    string
		arr1    []string
		arr2    []string
		results bool
	}{
		{name: "notok", arr1: []string{"a", "b", "c", "d", "e"}, arr2: []string{"m", "n", "o", "p"}, results: false},
		{name: "ok", arr1: []string{"a", "b", "c", "d", "e"}, arr2: []string{"m", "n", "o", "p", "e"}, results: true},
		{name: "emptyFirstSlice", arr1: []string{}, arr2: []string{"m", "n", "o", "p", "e"}, results: false},
		{name: "emptySecondSlice", arr1: []string{"a", "b", "c", "d", "e"}, arr2: []string{}, results: false},
	}
	for _, test := range testsCases {
		res := CompareIfItemExist2(test.arr1, test.arr2)
		if res != test.results {
			fmt.Printf("case %s failed it expected %t but got %t", test.name, test.results, res)
		}
	}
	fmt.Println("test passed succesifully")
}

// O(n*m) time complexity and space complexity O(1)
func CompareIfItemExist(arr1, arr2 []string) bool {
	var exist bool
	for _, v := range arr1 {
		for _, k := range arr2 {
			if v == k {
				exist = true
				break
			}
		}
	}
	return exist
}

// O(n+m) time complexity and space complexity O(n)
func CompareIfItemExist2(arr1, arr2 []string) bool {
	mapper := make(map[string]bool)
	for _, item := range arr1 {
		mapper[item] = true
	}
	var exist bool
	for _, item := range arr2 {
		if _, ok := mapper[item]; ok {
			exist = true
		}
	}
	return exist
}
