package main

import (
	"fmt"
	"sync"
)

func channelWaitgroup() {
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	// Add the number of works equal to the number of array elements.
	wg.Add(len(array))
	go func() {
		for _, value := range array {
			d := value
			ch <- &d
		}
	}()
	go func() {
		for value := range ch {
			fmt.Println(*value)
			// Decrement the waitgroup counter with each iteration.
			wg.Done()
		}
	}()

	wg.Wait()
}
func vamosChannel() {
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for _, value := range array {
			d := value
			ch <- &d
		}
		// Close channel to reach wg.Done() line.
		close(ch)
	}()
	go func() {
		for value := range ch {
			fmt.Println(*value)
		}
		wg.Done()
	}()
	wg.Wait()
}
