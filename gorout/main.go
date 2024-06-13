package main

import (
	"fmt"
	"sync"
)

func printSomething(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(str)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go printSomething("this is printed in the goroutine", &wg)
	wg.Wait()
}
