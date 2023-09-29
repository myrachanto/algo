package main

import (
	"fmt"
	"os"
)

func main() {
	var sum float64
	var count int
	for {
		var v float64
		if _, err := fmt.Fscanln(os.Stdin, &v); err != nil {
			break
		}
		sum += v
		count++
	}
	if count == 0 {
		fmt.Fprintln(os.Stdout, "There was no items written from the cli")
		return
	}
	fmt.Fprintln(os.Stdin, "the average is", sum/float64(count))
}
