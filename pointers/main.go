package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var p person = person{"tony", 32}
	fmt.Printf("p is of type %T its value is %#v its value is %v \n", p, p, p)
	p2 := person{"tony", 32}
	fmt.Printf("p2 is of type %T its value is %#v its value is %v\n", p2, p2, p2)
	p3 := &person{"tony", 32}
	fmt.Printf("p3 is of type %T its value is %#v its value is %v\n", p3, p3, p3)
	if p == p2 {
		fmt.Println("p is equal to p2")
	}
	if p == *p3 {
		fmt.Println("p is equal to p3")
	}
	var p4 struct {
		name string
		age  int
	} = struct {
		name string
		age  int
	}{name: "tony", age: 32}

	if p == p4 {
		fmt.Println("p is equal to p4")
	}
	fmt.Printf("p4 is of type %T its value is %#v its value is %v\n", p4, p4, p4)
	var p5 struct {
		age  int
		name string
	} = struct {
		age  int
		name string
	}{name: "tony", age: 32}

	if p == p5 {
		fmt.Println("p is equal to p4")
	}
	fmt.Printf("p5 is of type %T its value is %#v its value is %v\n", p5, p5, p5)
}
