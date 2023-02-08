package main

import (
	"fmt"
	"log"
	"math/rand"

	"golang.org/x/sync/errgroup"
)

func Task(task int) error {
	if rand.Intn(10) == task {
		return fmt.Errorf("Task %v failed ", task)
	}
	fmt.Printf("Task %v completed\n", task)
	return nil
}
func main() {
	eg := &errgroup.Group{}
	for i := 0; i < 10; i++ {
		task := i
		eg.Go(func() error {
			return Task(task)
		})
	}
	if err := eg.Wait(); err != nil {
		log.Fatal("Error", err)
	}
	fmt.Println("Completed successfully!")
}
