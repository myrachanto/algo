package main

import (
	"fmt"
	"sort"
	"strings"
)

type res struct{ k, v int }

func main() {
	// data1 := []int{1, 3, 4, 5}
	// res := missingNum(data1)
	// fmt.Println(res)
	numbers := []int{3, 3, 4, 2, 2, 2, 5, 5, 5, 5}
	majority := findMajorityElement(numbers)
	fmt.Printf("The majority element is: %d\n", majority)

	numbers2 := []int{11, 15, 7, 5, 4, 2, 9, 0}
	// numbers2 := []int{1, -5, -5, 0}
	target := 9

	result2 := findSumInts(target, numbers2)
	fmt.Printf("Indices of the two numbers: %v\n", result2)

	testCases1 := []string{"()", "()[]{}", "(]", "([)]", "{[]}", ""}

	for _, testCase := range testCases1 {
		isValid := BracketBalanced1(testCase)
		fmt.Printf("Is '%s' valid? %v\n", testCase, isValid)
	}
	res := reverseString("hello")
	fmt.Println("------------reversed", res)

	testCases := []string{"racecar", "hello", "A man a plan a canal Panama", "abba"}

	for _, testCase := range testCases {
		palindrome := isPalindrome(testCase)
		fmt.Printf("Is '%s' a palindrome? %v\n", testCase, palindrome)
	}
	sum := dubblenum(30)
	fmt.Println("------------totalsum", sum)

}

func isPalindrome(str string) bool {
	s := strings.ToLower(str)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
func dubblenum(num int) int {
	sum := 1
	for i := 1; i < num; i++ {
		sum = sum * 2
	}
	return sum
}
func doubler(val int, days int) int {
	if days == 1 {
		return 0
	}

	return val + doubler(val*2, days-1)

}
func reverseString(str string) string {
	res := strings.Split(str, "")

	final := make([]string, 0, len(res))
	for i := len(res) - 1; i > -1; i-- {
		final = append(final, res[i])
	}
	return strings.Join(final, "")
}

func mytargetSolution(target int, arr []int) []res {
	complemetary := make(map[int]int)
	found := make(map[int]bool)
	for _, v := range arr {
		if v <= target {
			complemetary[v] = target - v
		}
	}
	ress := []res{}
	for _, k := range arr {
		if r, ok := complemetary[k]; ok {
			if _, ok := complemetary[r]; ok {
				if k >= r && !found[k] {
					ress = append(ress, res{k, r})
					found[k] = true
				}
			}
		}
	}
	return ress
}
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

func twoSum(nums []int, target int) []int {
	numIndices := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, found := numIndices[complement]; found {
			return []int{index, i}
		}
		numIndices[num] = i
	}

	return []int{} // No solution found
}
func BracketBalanced(str string) bool {
	res := true
	compare := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
		'<': '>',
	}
	mapa := make([]rune, 0, len(str))
	for _, y := range str {
		if _, ok := compare[y]; ok {
			mapa = append(mapa, y)
		} else if len(mapa) == 0 {
			res = false
			break
		} else if compare[mapa[len(mapa)-1]] != y {
			res = false
			break
		} else {
			// fmt.Println(mapa)
			mapa = mapa[:len(mapa)-1]
		}
	}
	if len(mapa) != 0 {
		res = false
	}
	return res

}
func BracketBalanced1(str string) bool {
	res := true
	compare := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
		'<': '>',
	}
	stack := make([]rune, 0)

	for _, char := range str {
		if openingBracket, ok := compare[char]; ok {
			stack = append(stack, openingBracket)
		} else if len(stack) == 0 || char != stack[len(stack)-1] {
			res = false
			break
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return res && len(stack) == 0
}

func IsBalanced2(text string) bool {
	isBalanced := true
	bMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	s := make([]rune, 0, len(text))

	for _, c := range text {
		if _, ok := bMap[c]; ok {
			s = append(s, c)
		} else if len(s) == 0 {
			isBalanced = false
			break
		} else if bMap[s[len(s)-1]] != c {
			isBalanced = false
			break
		} else {
			s = s[:len(s)-1]
		}
	}

	if len(s) != 0 {
		isBalanced = false
	}

	return isBalanced
}
func missingNum(nums []int) int {
	n := len(nums)

	expectedSum := (n * (n + 1)) / 2

	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}
	missingNumber := expectedSum - actualSum
	return missingNumber
}

// func mapSlice(data map[int]int) []int{

// }

func findMajorityElement(nums []int) int {
	mapa := make(map[int]int)
	for _, v := range nums {
		mapa[v]++
	}
	type res struct{ key, value int }
	ress := []res{}
	for k, v := range mapa {
		ress = append(ress, res{k, v})
	}
	sort.Slice(ress, func(i, j int) bool {
		return ress[i].value > ress[j].value
	})
	return ress[0].key
}
