package main

import "fmt"

func main() {
	// channelWaitgroup()
	// vamosChannel()
	deadLockCheck()
}

func slices() {
	array := []int{1, 2, 3, 4}
	newArray := make([]*int, 4)
	for i, value := range array {
		d := value
		newArray[i] = &d
	}
	for _, value := range newArray {
		fmt.Println(*value)
	}
}
