package main

import (
	"os"
	"sync"
	"testing"
)

// Helper function to create a temporary JSON file for testing
func createTempFile(content string) (string, error) {
	file, err := os.CreateTemp("", "*.json")
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

// Test for successful file processing
func TestJsonReader_Success(t *testing.T) {
	content := `[{"name": "John", "age": 30, "role": 1}, {"name": "Jane", "age": 25, "role": 2}]`
	filename, err := createTempFile(content)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(filename)

	var wg sync.WaitGroup
	work := make(chan string, 1)
	results := make(chan Task, 10)
	errs := make(chan error, 1)

	wg.Add(1)
	go jsonReader(&wg, work, results, errs)
	work <- filename
	close(work)
	wg.Wait()
	close(results)
	close(errs)

	var tasks []Task
	for task := range results {
		tasks = append(tasks, task)
	}

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

	if len(errs) > 0 {
		t.Errorf("Unexpected error: %v", <-errs)
	}
}

// Test for file processing errors
func TestJsonReader_FileError(t *testing.T) {
	var wg sync.WaitGroup
	work := make(chan string, 1)
	results := make(chan Task, 10)
	errs := make(chan error, 1)

	wg.Add(1)
	go jsonReader(&wg, work, results, errs)
	work <- "nonexistent_file.json"
	close(work)
	wg.Wait()
	close(results)
	close(errs)

	if len(results) > 0 {
		t.Errorf("Unexpected result received")
	}

	if len(errs) == 0 {
		t.Errorf("Expected an error but did not receive one")
	}
}
func TestJsonReader_EmptyFile(t *testing.T) {
	filename, err := createTempFile("")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(filename)

	var wg sync.WaitGroup
	work := make(chan string, 1)
	results := make(chan Task, 10)
	errs := make(chan error, 1)

	wg.Add(1)
	go jsonReader(&wg, work, results, errs)
	work <- filename
	close(work)
	wg.Wait()
	close(results)
	close(errs)

	if len(results) > 0 {
		t.Errorf("Unexpected result received")
	}

	if len(errs) > 0 {
		t.Errorf("Unexpected error: %v", <-errs)
	}
}

// Test for concurrent processing
func TestJsonReader_ConcurrentProcessing(t *testing.T) {
	content1 := `[{"name": "John", "age": 30, "role": 1}]`
	content2 := `[{"name": "Jane", "age": 25, "role": 2}]`
	file1, err := createTempFile(content1)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file1)
	file2, err := createTempFile(content2)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file2)

	var wg sync.WaitGroup
	work := make(chan string, 2)
	results := make(chan Task, 10)
	errs := make(chan error, 2)

	wg.Add(2)
	go jsonReader(&wg, work, results, errs)
	go jsonReader(&wg, work, results, errs)
	work <- file1
	work <- file2
	close(work)
	wg.Wait()
	close(results)
	close(errs)

	var tasks []Task
	for task := range results {
		tasks = append(tasks, task)
	}

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}

	if len(errs) > 0 {
		t.Errorf("Unexpected error: %v", <-errs)
	}
}
