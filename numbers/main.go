package main

import (
	"fmt"
	"strings"
)

func main() {
	flo := 600
	// ////////////////////////////////////////////////
	fmt.Printf("%s \n", strings.Repeat("/", 50))

	fmt.Println("converting Int 500 to uint", uint64(flo))

	fmt.Printf("%s \n", strings.Repeat("/", 50))

	fmt.Println("converting Int 500 to uint", int32(flo))
	// var i int32 = 7
	// var j int64 = 8
	// k = i+int32(j)

	///////////////////////////////////////////////////
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1
	fmt.Println("=============================2", slice2)
	// slice3 := []int{}
	slice3 := make([]int, 3)
	copy(slice3, slice1)
	fmt.Println("=============================3", slice3)
}

// We have a company with many employees and the management decides to give them different kinds of bonuses, based on different rules. New bonuses can be introduced at any time.
// Examples of bonus:

// 5% salary bonus on the birthday month for all employees younger than 30

// $500 bonus for all women who have children, on the birthday of their children

// 10% bonus for regular employee, 15% for management employee paid for company
