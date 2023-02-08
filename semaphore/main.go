package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"

	"golang.org/x/sync/semaphore"
)

var (
	wg   sync.WaitGroup
	sema = semaphore.NewWeighted(12)
)

func dosomething(i int) {
	wg.Add(1)
	defer wg.Done()
	defer sema.Release(1)
	fmt.Printf("%d ======> is the number of Goroutines running \n", runtime.NumGoroutine())
	if i%5 == 0 {
		fmt.Printf("%d is divisble by five\n", i)
	}
	if i%7 == 0 {
		fmt.Printf("%d is divisble by seven\n", i)
	}
	if i%11 == 0 {
		fmt.Printf("%d is divisble by eleven\n", i)
	}
}
func main() {
	for i := 0; i < 400; i++ {
		if err := sema.Acquire(context.Background(), 1); err != nil {
			log.Fatal(err)
		}
		go dosomething(i)
	}
	wg.Wait()

}
