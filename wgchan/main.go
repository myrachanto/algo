package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(i int, resp chan any) {
	defer wg.Done()
	ti := i * int(time.Second)
	if i > 5 {
		ti = i / 2 * int(time.Second)
	}
	time.Sleep(time.Duration(ti))
	resp <- fmt.Sprintf("%d took %v time to completes its work", i, ti)
}
func main() {
	resp := make(chan any, 10)
	l := 10
	wg.Add(l)
	for i := 1; i <= l; i++ {
		go worker(i, resp)
	}
	wg.Wait()
	close(resp)
	for g := range resp {
		fmt.Println(g)
	}
}
