package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var (
	// wg sync.WaitGroup = sync.WaitGroup{}
	wg = &waitGroup{}
)

// naive implementation o wait groups
type waitGroup struct {
	counter int32
}

func (wg *waitGroup) Add(i int32) {
	atomic.AddInt32(&wg.counter, i)
}
func (wg *waitGroup) Wait() {
	for {
		if atomic.LoadInt32(&wg.counter) == 0 {
			return
		}
	}
}
func (wg *waitGroup) Done() {
	wg.Add(-1)
	if atomic.LoadInt32(&wg.counter) < 0 {
		panic("negative waig groups")
	}
}
func main() {
	goruoutines := 10
	wg.Add(int32(goruoutines))
	for i := 0; i < goruoutines; i++ {
		go worker(i)
	}
	wg.Wait()
}
func worker(i int) {
	defer wg.Done()
	time.Sleep(300 * time.Duration(i) * time.Microsecond)
	fmt.Printf("Goroutine %d took  %d \n", i, 300*time.Duration(i)*time.Microsecond)
}
