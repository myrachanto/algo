package main

import (
	"fmt"
	"time"
)

func main() {
	list := map[string]int{"tony": 32, "jenny": 25, "mark": 27, "carolina": 27, "alex": 32}
	for k, v := range list {
		list[k] = v + 2
	}
	fmt.Println("+++++++++++++", list)
	// value := 10
	for i := range 10 {
		fmt.Println(10 - i)
	}
	for i := range 10 {
		go func() {
			fmt.Println("he value of i", i)
		}()
	}
	time.Sleep(1 * time.Second)
}
