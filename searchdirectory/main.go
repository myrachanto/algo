package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func searchInDirectory(root string, filename string, results chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if !info.IsDir() && info.Name() == filename {
			results <- path
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <directory> <filename>")
		os.Exit(1)
	}

	root := os.Args[1]
	filename := os.Args[2]

	results := make(chan string, 100)
	var wg sync.WaitGroup

	// Launch a goroutine for each subdirectory
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if info.IsDir() {
			wg.Add(1)
			go searchInDirectory(path, filename, results, &wg)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Close the results channel after all goroutines have finished
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect and print results
	for result := range results {
		fmt.Println("Found:", result)
	}
}