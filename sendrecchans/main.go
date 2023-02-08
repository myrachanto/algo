package main

import "fmt"

func main() {
	c := make(chan int, 4)
	producer(c)
	consumer(c)
}
func producer(c chan<- int) {
	c <- 4
	c <- 5
	c <- 6
	c <- 7
	close(c)
	// <-c
	// badfunc(c)
}
func consumer(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
	// c<-8
}

// func badfunc(c chan int) {
// 	v := <-c

// }
