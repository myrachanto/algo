package main

import (
	"context"
	"fmt"
	"sync"
)

type result struct {
	goroutineId int
	key         int
	value       int
}

func work(ctx context.Context, wg *sync.WaitGroup, data []int, target int, goroutineId int, res chan<- result) {
	defer wg.Done()
	fmt.Printf("Goroutine %d started\n", goroutineId)
	for k, v := range data {
		select {
		case <-ctx.Done():
			// Log when a goroutine is canceled and exits
			fmt.Printf("Goroutine %d canceled\n", goroutineId)
			return
		default:
			if v == target {
				// Send the result to the channel and return immediately
				select {
				case res <- result{goroutineId: goroutineId, key: k, value: v}:
					fmt.Printf("Goroutine %d found the target at index %d with value %d\n", goroutineId, k, v)
					return
				case <-ctx.Done():
					fmt.Printf("Goroutine %d canceled\n", goroutineId)
					return
				}
			}
		}
	}
	// Log completion if the target is not found and no cancellation signal is received
	fmt.Printf("Goroutine %d completed without finding the target\n", goroutineId)
}
func generateData(limit int) []int {
	res := make([]int, limit)
	for i := range limit {
		res[i] = i
	}
	return res
}
func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	dataSet := generateData(1000000)
	start, limit, numberOfGoroutines, target := 0, 100000, 10, 300234
	wg.Add(numberOfGoroutines)
	results := make(chan result, 1)
	for i := range numberOfGoroutines {
		end := start + limit
		if end > len(dataSet) {
			end = len(dataSet) // Avoid out-of-bounds slice access
		}
		go work(ctx, &wg, dataSet[start:end], target, i, results)
		start = end
	}
	// Goroutine to close results channel after all goroutines are done
	go func() {
		wg.Wait()
		close(results)
	}()
	// Read from results channel
	found := false
	for res := range results {
		fmt.Printf("Main received result from Goroutine %d: target found at index %d with value %d\n", res.goroutineId, res.key, res.value)
		cancel() // Cancel context to stop other goroutines
		found = true
		break
	}

	if !found {
		fmt.Println("No result found")
	}

}
