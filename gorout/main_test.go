package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printsomething(t *testing.T) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	var wg sync.WaitGroup
	wg.Add(1)
	go printSomething("there you go", &wg)
	wg.Wait()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdout
	if !strings.Contains(output, "there you go") {
		t.Errorf("expected there you go but its not there")
	}
}
