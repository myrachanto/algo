package main

import "fmt"

func firstRepeatableNum(data []int) int {
	mapper := make(map[int]bool)
	res := -1
	for _, k := range data {
		if _, ok := mapper[k]; ok {
			res = k
			break
		} else {
			mapper[k] = true
		}
	}
	return res
}
func main() {
	testCases := []struct {
		name string
		data []int
		res  int
	}{
		{"ok", []int{3, 4, 5, 4, 5, 6, 7, 8, 7, 5, 6}, 4},
		{"notok", []int{1, 2, 3, 4, 56, 7, 8, 9}, -1},
	}
	for _, test := range testCases {
		res := firstRepeatableNum(test.data)
		if res != test.res {
			fmt.Printf(" expected %d but got %d \n", res, test.res)
		} else {
			fmt.Printf("%s passed \n", test.name)
		}
	}
}
