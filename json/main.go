package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	data := map[string]interface{}{
		"name":  "Jane Doe",
		"email": "jane@example.com",
	}
	jsonWriter(data)
	jsonWriterAdvanced(data)

}
func jsonWriter(data map[string]interface{}) error {
	// Create or overwrite the file
	file, err := os.Create("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a JSON encoder
	encoder := json.NewEncoder(file)

	// Encode and write the JSON data
	err = encoder.Encode(data)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
func jsonWriterAdvanced(data map[string]interface{}) error {
	// Create or overwrite the file
	file, err := os.OpenFile("dataAdvanced.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a JSON encoder
	encoder := json.NewEncoder(file)

	// Encode and write the JSON data
	err = encoder.Encode(data)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
