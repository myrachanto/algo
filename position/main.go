package main

import (
	"fmt"
	"time"
)

var (
	mapper  = []pos{{"start", 0, 0}}
	counter = 0
	res     = false
)

type pos struct {
	direction string
	length    int
	offset    int
}

func findAlgorithm(d, offset int, data []int) {
	counter++
	l := len(data)
	if l == 1 && data[0] == d {
		fmt.Println(">>>>>>>>>>>>>>", data[0], l, d)
		res = true
		fmt.Println(">>>>>>>>>>>>>>", data[0], l, d, res)
		return
	}
	if l == 1 && data[0] != d {
		return
	}
	half := l / 2
	left, right := data[:half], data[half:]
	if left[len(left)-1] >= d {
		pos := pos{"left", counter, 0}
		mapper = append(mapper, pos)
		findAlgorithm(d, 0, left)
	} else {
		lenright := len(left)
		pos := pos{"right", counter, lenright}
		mapper = append(mapper, pos)
		findAlgorithm(d, lenright, right)
	}
}
func getsum() (res int) {
	for _, v := range mapper {
		res += v.offset
	}
	return res
}
func generator() []int {
	results := []int{}
	for i := 1; i <= 20000; i++ {
		if i%2 == 0 {
			results = append(results, i)
		}
	}
	return results
}
func main() {
	num := 6000
	data := generator()
	start := time.Now()
	position := -1
	findAlgorithm(num, 0, data)
	if res {
		position = getsum()
	}
	fmt.Println(mapper)
	fmt.Println(position)
	fmt.Println(time.Since(start))
}

func best(d int, data []int, ch chan<- int) {
	res := -1
	for k, v := range data {
		if v == d {
			ch <- k
			break
		}
		if v > d {
			break
		}
	}
	ch <- res
}

func sum() int {
	position, left, right := 0, 0, 0
	for k, v := range mapper {
		if k == 1 && v.direction == "right" {
			position = v.length
		} else if k != 1 && v.direction == "right" {
			right += v.length
		} else {
			left += v.length
		}
	}
	fmt.Println("------------", position, left, right)
	if right == 0 {
		return position
	} else {
		position = position + right - left
		return position
	}
}
