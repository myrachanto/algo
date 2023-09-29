package main

import (
	"fmt"
	"time"
)

type results struct {
	name string
}

func main() {
	done := make(chan results)
	go func() {
		work()
		done <- results{"tony"}
	}()
	res := <-done

	fmt.Println("....closing time", res)
}
func work() {
	time.Sleep(1 * time.Millisecond)
	fmt.Println("work do doing its job")
}
