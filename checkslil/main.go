package main

import "fmt"

func main() {
	mapper := []int{1, 2, 3, 4, 5}
	target := 7

	res, ok := EvaluateIfSecionExist(target, mapper)
	fmt.Println("==============", res, ok)
	if !ok {
		fmt.Println("==============", res)
	}

}
func EvaluateIfSecionExist(code int, newInfo []int) (int, bool) {
	var results bool = false
	for _, g := range newInfo {
		if g == code {
			results = true
			break
		}
	}
	return code, results
}
