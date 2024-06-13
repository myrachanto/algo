package main

import "fmt"

func findTargetSum(nums []int, target int) []int {
	var result []int
	backtrack(nums, target, 0, 0, &result)
	return result
}

func backtrack(nums []int, target, currentSum, startIndex int, result *[]int) {
	if currentSum == target {
		// Found a valid combination
		*result = append(*result, nums[startIndex:]...)
		return
	}

	for i := startIndex; i < len(nums); i++ {
		currentSum += nums[i]
		*result = append(*result, nums[i])

		// Explore further with the updated currentSum and startIndex
		backtrack(nums, target, currentSum, i+1, result)

		// Backtrack by removing the last element
		currentSum -= nums[i]
		*result = (*result)[:len(*result)-1]
	}
}

func main() {
	nums := []int{2, 3, 7, 4, 5}
	target := 10

	result := findTargetSum(nums, target)
	if len(result) > 0 {
		fmt.Printf("Subset with target sum %d: %v\n", target, result)
	} else {
		fmt.Printf("No subset found with target sum %d\n", target)
	}
	res := findSumInts(target, nums)
	fmt.Println(res)
}

type res struct{ k, v int }

func findSumInts(target int, arr []int) []res { 
	ress := []res{}
	found := make(map[int]bool)
	for _, v := range arr {
		complement := target - v
		if found[complement] {
			ress = append(ress, res{v, complement})
		}
		found[v] = true
	}
	return ress

}
func TwoSum(numbers []int, target int) [2]int {
	mapper := make(map[int]int)
	var res [2]int
	for i, num := range numbers {
		complement := target - num
		if idx, ok := mapper[complement]; ok {
			res[0] = idx
			res[1] = i
			return res
		}
		mapper[num] = i
	}
	return res
}
