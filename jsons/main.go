package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

	p2 := Person{Name: "Alice", Age: 30}
	file, err := os.Create("person.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(&p2); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}
