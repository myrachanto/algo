package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	raw := []byte("i love golang")
	if n, err := handleData(bytes.NewReader(raw)); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Bytes read: %d\n", n)
	}

	// Example with a file
	file, err := os.Open("./interfaces2/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if n, err := handleData(file); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Bytes read: %d\n", n)
	}

	if n, err := handleData2(raw); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Bytes read: %d\n", n)
	}
}
func handleData(r io.Reader) (int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	return len(data), nil
}
func handleData2(r []byte) (int, error) {
	return len(r), nil
}
