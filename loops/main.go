package main

import (
	"fmt"
	"slices"
)

func main() {
	// golang 1.22 is here with new features

	// ranging over numbers
	for i := range 7 {
		fmt.Println(i)
	}
	// concatinating two slices of the same type
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5}
	arr3 := slices.Concat(arr1, arr2)
	fmt.Println(arr3)
	
	// The long-standing “for” loop gotcha with accidental sharing of loop variables between iterations is now resolved. Starting with Go 1.22, the following code will print “a”, “b”, and “c” in some order:

    done := make(chan bool)

    values := []string{"a", "b", "c"}
    for _, v := range values {
        go func() {
            fmt.Println(v)
            done <- true
        }()
    }

    // wait for all goroutines to complete before exiting
    for _ = range values {
        <-done
    }

}
