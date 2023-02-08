package main

import "fmt"

func main() {
	// Two was to initialize slices
	s1 := make([]string, 5) //this gives you the power to give a slice length
	s2 := []string{"hi"}    // this gives you the power to initialize slices with values
	fmt.Println(s1, s2)
	// declare slices
	var s3 []string // this is slice declaration and the slice is nil
	//a slice is a window to an array
	if s3 == nil {
		fmt.Printf("s3 -> %v \n", s3)
	}
	a := [5]int{1, 2, 3, 4, 5}
	s4 := a[:3]
	fmt.Printf("s4 contains %v and cap is %d and len is %d \n", s4, cap(s4), len(s4)) // [1,2,3]
	s5 := s4[:4]
	fmt.Printf("s5 contains %v and cap is %d and len is %d \n", s5, cap(s5), len(s5)) // [1,2,3,4]
	s6 := a[0:3:3]
	fmt.Printf("s6 contains %v and cap is %d and len is %d \n", s6, cap(s6), len(s6)) // [1,2,3]
	// s7 := s6[:4]
	// fmt.Printf("s5 contains %v and cap is %d and len is %d \n", s7, cap(s7), len(s7)) // this errors out coz the capacity of s7 is 3
	// Append
	s4 = append(s4, 7,8,9,10,11)
	fmt.Printf("s4 contains %v and cap is %d and len is %d \n", s4, cap(s4), len(s4)) // [1,2,3,7]

	fmt.Printf("s5 contains %v and cap is %d and len is %d \n", s5, cap(s5), len(s5)) // [1,2,3,7]
	s6 = append(s6, 6, 7)
	fmt.Printf("s6 contains %v and cap is %d and len is %d \n", s6, cap(s6), len(s6)) // [1,2,3,6]
	s7 := make([]int, 5,5)
	s7 = append(s7, 2,3) // [0,0,0,0,0,2,3]
	fmt.Printf("s7 contains %v and cap is %d and len is %d \n", s7, cap(s7), len(s7)) // [1,2,3,6]
	// filter out an element in a slice
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sl2 := []int{}
	for _, j := range sl {
		if j == 5 {
			continue
		}
		sl2 = append(sl2, j)
	}
	fmt.Println(sl2)
	copy(sl,s6)
	fmt.Println(sl)
}
