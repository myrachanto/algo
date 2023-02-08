package main

import "fmt"

type soma float64

func (s soma) String() string {
	return fmt.Sprintln("me checking the magic of stringer method in golang")
}
func main() {
	var s soma
	fmt.Println(s)
}
