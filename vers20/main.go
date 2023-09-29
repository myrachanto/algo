package main

import "fmt"

type resa struct {
	num       int
	numindex  int
	comp      int
	compindex int
}
type results struct {
	resa []resa
}

func (r *results) adder(data []int, target int) {
	visited := make(map[int]int)
	for k, v := range data {
		x, ok := checkViablity(v, target)
		if ok {
			res, ok := visited[x]
			if ok {
				r.resa = append(r.resa, resa{v, k, x, res})
			}
			visited[v] = k
		}
	}
}
func checkViablity(num, target int) (int, bool) {
	return target - num, num < target
}

func (s *results) Printer() {
	for _, v := range s.resa {
		fmt.Printf("number %d of index %d is a complement to %d of index %d\n", v.num, v.numindex, v.comp, v.compindex)
	}
}
func main() {
	// target 9 get two numbers that add to 9
	arr := []int{2, 11, 6, 15, 3, 7}
	target := 9
	results := results{}
	// res := adder1(arr, target)
	// fmt.Printf("%#v", res)
	results.adder(arr, target)
	results.Printer()
}

func adder1(data []int, target int) []resa {
	visited := make(map[int]int)
	results := []resa{}
	for k, v := range data {
		ok := checkIfValid(v, target)
		if ok {
			x := checkComplementary(v, target)
			res, ok := visited[x]
			if ok {
				results = append(results, resa{v, k, x, res})
			}
			visited[v] = k
		}
	}
	return results
}
func checkComplementary(num, target int) int {
	return target - num
}
func checkIfValid(num, target int) bool {
	return num < target
}
