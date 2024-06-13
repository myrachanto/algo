package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// str := "codewars"
	// fmt.Println(str[1:5])
	// red := GetCount("this is cool")
	// fmt.Println("-----------", red)
	// res := MinMax([]int{1, 2, 3, 4, 5})
	// fmt.Println("-----------", res)
	res := CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	fmt.Println(res)
	

}
func CreatePhoneNumber(numbers [10]uint) string {

	ms := numbers
	res := fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", ms[0], ms[1], ms[2], ms[3], ms[4], ms[5], ms[6], ms[7], ms[8], ms[9])
	return res
} 
func Pyramid(n int) [][]int {
	// your code here
	pyramid := make([][]int, 0, n)
	for i := 0; i < n; i++ {
		box := []int{}
		if i == 0 {
			box = append(box, 0)
		} else {
			for j := 0; j <= i; j++ {
				box = append(box, 1)
			}
		}
		pyramid = append(pyramid, box)
	}
	return pyramid
}
func TowerBuilder(nFloors int) []string {
	liststars := []int{1}
	for i := 1; i < nFloors; i++ {
		next := liststars[len(liststars)-1] + 2
		liststars = append(liststars, next)
	}
	res := []string{}
	base := liststars[len(liststars)-1]
	for _, k := range liststars {
		base1 := base - k
		b := base1 / 2
		res = append(res, fmt.Sprintf("%s%s%s\n", strings.Repeat(" ", b), strings.Repeat("*", k), strings.Repeat(" ", b)))
	}
	return res
}
func MinMax(arr []int) [2]int {
	min, max := getMinMax(arr)
	fmt.Println("------", min, max)
	res := [2]int{min, max}
	return res
}
func getMinMax(d []int) (int, int) {
	min := d[0]
	max := d[0]
	for _, g := range d {
		if min > g {
			min = g
		}
		if max < g {
			max = g
		}
	}
	return min, max
}
func GetCount1(str string) (count int) {
	for _, c := range str {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
func GetCount(str string) (count int) {
	// Enter solution here
	mapper := make(map[rune]int)
	exist := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	for _, bs := range str {
		if _, ok := exist[bs]; ok {
			mapper[bs]++
		}
	}
	sum := 0
	for _, k := range mapper {
		sum += k
	}
	return sum
}

func ReverseWords(str string) string {
	res0 := strings.Split(str, " ")
	results := []string{}
	for _, word := range res0 {
		reversed := []string{}
		wordchar := strings.Split(word, "")
		for i := len(wordchar) - 1; i >= 0; i-- {
			reversed = append(reversed, wordchar[i])
		}
		results = append(results, strings.Join(reversed, ""))
	}
	return strings.Join(results, " ") // reverse those words
}
func CleanString(s string) string {
	stack := []rune{}
	for _, r := range s {
		if r == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, r)
		}
	}
	return string(stack)
}
func FindEvenIndex(arr []int) int {
	var sum, b int
	for _, n := range arr {
		sum += n
	}
	for i, n := range arr {
		sum -= n
		if sum == b {
			return i
		}
		b += n
	}
	return -1
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

func IsValidWalk(walk []rune) bool {
	if time(walk) != 10 {
		return false
	}

	p := &Position{}
	for _, direction := range walk {
		p.Move(direction)
	}
	return p.Equal(Position{})
}

func time(walk []rune) int {
	return len(walk)
}

type Position struct {
	x int
	y int
}

func (p *Position) Move(direction rune) {
	switch direction {
	case 'n':
		p.y++
	case 's':
		p.y--
	case 'w':
		p.x--
	case 'e':
		p.x++
	}
}
func (p Position) Equal(p2 Position) bool {
	return p.x == p2.x && p.y == p2.y
}

func SpinWords(str string) string {
	res := strings.Split(str, " ")
	results := []string{}
	for _, part := range res {
		if len(part) >= 5 {
			parta := strings.Split(part, "")
			parta2 := []string{}
			for i := len(parta) - 1; i >= 0; i-- {
				parta2 = append(parta2, parta[i])
			}
			part = strings.Join(parta2, "")
		}
		results = append(results, part)
	}
	return strings.Join(results, " ")
} // SpinWords
func FindOdd(seq []int) int {
	// your code here
	res := make(map[int]int)
	for _, g := range seq {
		res[g]++
	}
	num := -1
	for _, appearance := range res {
		if appearance%2 == 1 {
			num = appearance
			break
		}
	}
	return num
}
func FindOdd1(seq []int) (res int) {
	for _, n := range seq {
		res ^= n
	}
	return
}
func FoldArray(arr []int, runs int) []int {
	// Base case: If runs is 0 or the array is empty, return the original array.
	if runs <= 0 || len(arr) == 0 {
		return arr
	}

	// Create a new array to store the folded result.
	var result []int

	// Calculate the midpoint of the array.
	midpoint := len(arr) / 2

	// Iterate through the first half of the array and add the corresponding elements
	// from the second half by folding them.
	for i := 0; i < midpoint; i++ {
		result = append(result, arr[i]+arr[len(arr)-1-i])
	}

	// If the length of the original array is odd, add the middle element to the result.
	if len(arr)%2 == 1 {
		result = append(result, arr[midpoint])
	}

	// Recursively fold the result array for (runs-1) times.
	return FoldArray(result, runs-1)
}
func Order(sentence string) string {
	//   split the sentence
	//   get the runes of numbers and compare where they appear in the sentence
	//   the sort the word in regards to rune
	res := strings.Split(sentence, " ")
	results := []string{}
	for i := 1; i <= 9; i++ {
		for _, part := range res {
			if strings.Contains(part, fmt.Sprintf("%d", i)) {
				results = append(results, part)
			}
		}
	}

	return strings.Join(results, " ")
}

func testValidIp() {

	testCases := []struct {
		name  string
		input string
		res   bool
	}{
		{"ok", "123.200.3.1", true},
		{"notok", "03.456.1", false},
		{"notok2", "123.03.786.1", false},
		{"notok3", "123.03.254.256", false},
	}
	for _, test := range testCases {
		res := Is_valid_ip(test.input)
		if res != test.res {
			fmt.Printf("Expected %t but got %t for %s \n", res, test.res, test.name)
		} else {

			fmt.Printf("%s passed \n", test.name)
		}
	}
}

func Is_valid_ip(ip string) bool {
	res := net.ParseIP(ip)
	return res != nil
}
