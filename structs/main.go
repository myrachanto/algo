package main

import (
	"fmt"
	"log"
)

type Person struct {
	Name, Address string
	Age           uint8
}

func main() {
	p := new(Person)

	fmt.Println(p)
	// m1 := [3]int{}
	// m2 := [5]int{1, 2, 3, 4, 5}
	// m1 = m2 //must be of equal length
	ms1 := []int{}
	ms2 := []int{1, 2, 3, 4, 5}
	copy(ms2, ms1)
	defer writersomething()
	// panic("wheeeeeeeeeee")
	// fmt.Println("continues")
	log.Panicln("here we exit!")
	// fmt.Println("---------------")
}
func writersomething() {
	fmt.Println("vamos a bialas")
}
