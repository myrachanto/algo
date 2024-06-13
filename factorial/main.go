package main

import "fmt"

// recursiveFactorial(int) int // 5! = 5*4*3*2*1
func recursiveFactorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * recursiveFactorial(num-1)
}
// RecursiveFibinacci(int) int // 0, 1,1,2,3,5
func RecursiveFibinacci(num int) int {
	if num < 2 {
		return num // 0, 1
	}
	return RecursiveFibinacci(num-1) + RecursiveFibinacci(num-2)
}

// initialize a map to store the results
var mapper = make(map[int]int)
// RecursiveFibinacciEnhanced(int) int // 0, 1,1,2,3,5
func RecursiveFibinacciEnhanced(num int) int {
	if num < 2 {
		return num // 0, 1
	}
	// check if the result is already in the map
	if res, ok := mapper[num]; ok {
		return res // return the result
	}
	// if not in the map, calculate the result
	res := RecursiveFibinacci(num-1) + RecursiveFibinacci(num-2)
	// store the result in the map
	mapper[num] = res
	return res
}
// fibinacciIterative(int) int // 0, 1,1,2,3,5
func fibinacciIterative(num int) int {
	// initialize a slice to store the results and store the first two results
	results := []int{0, 1}
	// loop from 2 to num
	for i := 2; i <= num; i++ {
		// append the sum of the previous two results
		results = append(results, results[i-2]+results[i-1])
	}
	// return the last result
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
