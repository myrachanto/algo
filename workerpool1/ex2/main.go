package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Task struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	Role int    `json:"role,omitempty"`
}

func jsonReader(wg *sync.WaitGroup, work chan string, results chan Task, errs chan error) {
	defer wg.Done()
	for filename := range work {
		file, err := os.Open(filename) // Change the filename as needed
		if err != nil {
			errs <- err
			continue
		}
		// Ensure the file is closed after processing
		// Move this line right after opening the file
		func() {
			defer file.Close()

			decoder := json.NewDecoder(file)
			var tasks []Task
			if err := decoder.Decode(&tasks); err != nil {
				errs <- err
				return
			}
			for _, task := range tasks {
				results <- task
			}
		}()
	}

}

func main() {
	var wg sync.WaitGroup
	files := []string{"./workerpool1/ex2/data/file.json", "./workerpool1/ex2/data/file2.json", "./workerpool1/ex2/data/file3.json", "/workerpool1/ex2/data/file4.json"}
	workerpool := 2
	errs := make(chan error, len(files))
	work := make(chan string, len(files))
	results := make(chan Task, 50)
	// start the workers
	for i := 1; i <= workerpool; i++ {
		wg.Add(1)
		go jsonReader(&wg, work, results, errs)
	}
	//file the work
	for _, filenanme := range files {
		work <- filenanme
	}
	close(work)
	wg.Wait()
	res := []Task{}
	// Process results
	go func() {
		wg.Wait()
		close(results)
		close(errs) // Close the results channel after all workers finish
	}()

	for task := range results {
		res = append(res, task)
	}
	// Process errors
	for err := range errs {
		fmt.Println(err)
	}
	fmt.Println(len(res))

}
