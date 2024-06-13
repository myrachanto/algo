package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func listing1() {
	i := 0

	go func() {
		defer wg.Done()
		i++
	}()

	go func() {
		defer wg.Done()
		i++
	}()
	fmt.Println("listing 1-------", i)

}

func listing2() {
	var i int64

	go func() {
		defer wg.Done()
		atomic.AddInt64(&i, 1)
	}()

	go func() {
		defer wg.Done()
		atomic.AddInt64(&i, 1)
	}()
	fmt.Println("listing 2-------", i)
}

func listing3() {
	i := 0
	mutex := sync.Mutex{}

	go func() {
		defer wg.Done()
		mutex.Lock()
		i++
		mutex.Unlock()
	}()

	go func() {
		defer wg.Done()
		mutex.Lock()
		i++
		mutex.Unlock()
	}()
	fmt.Println("listing 3-------", i)
}

func listing4() {
	i := 0
	ch := make(chan int)

	go func() {
		defer wg.Done()
		ch <- 1
	}()

	go func() {
		defer wg.Done()
		ch <- 1
	}()

	i += <-ch
	i += <-ch
	fmt.Println("listing 4-------", i)
}

func listing5() {
	i := 0
	mutex := sync.Mutex{}

	go func() {
		defer wg.Done()
		mutex.Lock()
		defer mutex.Unlock()
		i = 1
	}()

	go func() {
		defer wg.Done()
		mutex.Lock()
		defer mutex.Unlock()
		i = 2
	}()

	_ = i
	fmt.Println("listing 5-------", i)
}
func main() {
	wg.Add(10)
	listing1()
	listing2()
	listing3()
	listing4()
	listing5()
	wg.Wait()
}
