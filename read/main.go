package main

import "sync"

var mu sync.RWMutex
var savings int

func readData() int {
	mu.RLock()
	defer mu.RUnlock()
	return savings
}

func deposit(val int) {
	mu.Lock()
	defer mu.Unlock()
	savings += val
}
func withdraw(val int) {
	mu.Lock()
	defer mu.Unlock()
	savings += val
}

func main() {
	// members := []string{"antony", "jenny", "margaret", "alex", "caroline"}
	// var wg sync.WaitGroup

}
