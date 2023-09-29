package main

import "fmt"

type pair struct {
	A, B int
}

func main() {
	testsCases := []struct {
		name    string
		data    []int
		target  int
		results []int
		status  bool
	}{
		{"ok", []int{3, 4, 6, 4, 5}, 8, []int{4, 4}, true},
		{"notok", []int{3, 4, 6, 4}, 11, []int{}, false},
		{"notok", []int{3, 4, 6, 4, 5}, 10, []int{}, true},
	}
	for _, test := range testsCases {
		res, ok := pairssum(test.data, test.target)
		if ok != test.status {
			fmt.Printf("%s was expecting %t but got %t \n", test.name, test.status, ok)
		} else {
			fmt.Printf("%v sum pairs is %v\n", test.name, res)
		}
	}
}
func pairssum(data []int, target int) ([]pair, bool) {
	found := make(map[int]bool)
	pairs := []pair{}
	for _, num := range data {
		complement := target - num
		if found[complement] {
			pairs = append(pairs, pair{num, complement})
		}
		found[num] = true
	}
	return pairs, len(pairs) > 0
}
