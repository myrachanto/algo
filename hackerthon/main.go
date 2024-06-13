package main

import (
	"fmt"
	"sync"
)

func fizzBuzz(n int32) {
	// Write your code here
	f := int(n)
	for i := 1; i <= f; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
func countBits(n int32) int {
	count := 0

	for i := 0; i < 32; i++ {
		// Check if the least significant bit is set
		if n&1 == 1 {
			count++
		}

		// Right shift the number to move to the next bit
		n >>= 1
	}

	return count
}

type server struct {
	in       []int32
	oddChan  chan int32
	evenChan chan int32
}

var (
	serverChan = make(chan server)
	wg         sync.WaitGroup
)

func Server() {
	defer wg.Done()

	for {
		select {
		case s, ok := <-serverChan:
			if !ok {
				return
			}
			if s.in[0]%2 == 0 {
				s.evenChan <- s.in[0]
			} else {
				s.oddChan <- s.in[0]
			}
		}
	}

}

func main() {
	// fizzBuzz(15)
	// s.in[0]%2
	number := int32(42)
	bitCount := countBits(number)
	fmt.Printf("Number of set bits in %d: %d\n", number, bitCount)
}
