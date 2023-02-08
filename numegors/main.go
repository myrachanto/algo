package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func produce(data []int, done chan bool, hay int, position chan int, offset int) {
	defer wg.Done()
	for k, v := range data {
		if v == hay {
			done <- true
			position <- offset + k
			break
		}
	}
}
func generate() (res []int) {
	for i := 0; i < 100000; i++ {
		res = append(res, i)
	}
	return
}
func main() {
	statuschan := make(chan bool, 1)
	position := make(chan int, 1)
	data := generate()
	hay := 26001
	wg.Add(10)
	start := time.Now()
	for i := 0; i < 10; i++ {
		v := i * 10000
		fmt.Printf("%d are the running goroutines \n", runtime.NumGoroutine())
		go produce(data[v:v+10000], statuschan, hay, position, v)
	}
	wg.Wait()
	end := time.Since(start)
	close(statuschan)
	close(position)
	fmt.Printf("%s\n", strings.Repeat("/", 50))
	fmt.Printf("%d are the running goroutines\n", runtime.NumGoroutine())
	fmt.Printf("%d exist \"%v\" at position %d \n", hay, <-statuschan, <-position)
	fmt.Println("Time taken", end)

	fmt.Printf("%s\n", strings.Repeat("/", 50))

}
