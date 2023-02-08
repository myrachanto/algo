package main

import "fmt"

func main(){
	// make() and new()
	// built -n functions to allocate memory for a given target type T
	//  make(t ,args) and new(T)

	// DIFFRERENCES 
	// data types they work with
	// make( ) only works with slices, maps and channels
	_ = make(map[int]int)
	// or
	_ = make([]int, 7)
	// or
	_ = make(chan int)
	// new() works with any
	_ = new(int)
	// or 
	_ = new(map[int]int)

	// RETURNED TYPES
	// make() returns the target type 
	// new() returns the address
	var v map[int]int = map[int]int{}
	var p *map[int]int = new(map[int]int)
	fmt.Println(p,v)

	// MEMORY INITIALIZATION
	// make() initializes memory
	// new() zeros the memory or nil
}