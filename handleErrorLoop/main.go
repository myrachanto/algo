package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	list2 := []int{300, 400, 500}
	results, err := checkJakport(list, list2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("results", results)
}

func checkJakport(list1, list2 []int) ([]int, error) {
	var results []int
	for _, v := range list1 {
		for _, k := range list2 {
			res := k / v
			results = append(results, res)
			for _, r := range results {
				final := r % 10
				if final > 0 {
					return results, fmt.Errorf("hit the low level")
				}
			}
		}
	}
	return results, nil
}
