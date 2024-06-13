package main

import (
	"fmt"
	"sync"
)

func deadLockCheck() {
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for _, value := range array {
			ch <- &value
		}
	}()
	go func() {
		for value := range ch {
			fmt.Println(*value)
		}
		wg.Done()
	}()

	// New goroutine is run.
	go func() {
		var a int
		for {
			a++
			fmt.Println(a)
		}
	}()

	wg.Wait()
}
