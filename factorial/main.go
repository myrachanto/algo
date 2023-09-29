package main

import "fmt"

func recursiveFactorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * recursiveFactorial(num-1)
}

func RecursiveFibinacci(num int) int {
	if num < 2 {
		return num
	}
	return RecursiveFibinacci(num-1) + RecursiveFibinacci(num-2)
}

var mapper = make(map[int]int)

func RecursiveFibinacciEnhanced(num int) int {
	if num < 2 {
		return num
	}
	if res, ok := mapper[num]; ok {
		return res
	}
	res := RecursiveFibinacci(num-1) + RecursiveFibinacci(num-2)
	mapper[num] = res
	return res
}
func fibinacciIterative(num int) int {
	results := []int{0, 1}
	for i := 2; i <= num; i++ {
		results = append(results, results[i-2]+results[i-1])
	}
	return results[num]
}

func main() {
	res := recursiveFactorial(5)
	//   5*4*3*2*1 = 120
	fmt.Println(res)

	res = RecursiveFibinacci(5)
	// 0, 1,1,2,3,5
	fmt.Println(res)
	res = fibinacciIterative(5)
	// 0, 1,1,2,3,5
	fmt.Println(res)
}
