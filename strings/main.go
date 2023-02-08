package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	data := `{"name": "tony", "age": 35, "email": "example@gmail.com"}`
	type user struct {
		Name  string `json:"name,omitempty"`
		Age   uint8  `json:"age,omitempty"`
		Email string `json:"email,omitempty"`
	}
	ok := json.Valid([]byte(data))
	if !ok {
		log.Fatal("your json data object is invalid")
		return
	}
	var u user
	err := json.Unmarshal([]byte(data), &u)
	if err != nil {
		log.Fatal("something went wrong marshaling json data")
		return
	}
	// err := json.NewDecoder().Decode(u)
	// if err != nil {
	// 	log.Fatal("something went wrong marshaling json data")
	// 	return
	// }
	fmt.Println("---------------------------")
	fmt.Printf("%v \n", u)
	fmt.Println("---------------------------")
	jsondata1, err := json.MarshalIndent(u, "", "\t")
	if err != nil {
		log.Fatal("something went wrong marshaling json data")
		return
	}
	fmt.Println("---------------------------")
	fmt.Printf("%s \n", jsondata1)
	fmt.Printf("%v \n", string(jsondata1))
	response, err := http.Post("http://task1:2100/api/task1/", "application/json", bytes.NewBuffer(jsondata1))
	if err != nil {
		log.Fatal("something went wrong with the post request")
		return
	}
	var u2 user
	err = json.NewDecoder(response.Body).Decode(&u2)
	if err != nil {
		log.Fatal("something went wrong deconding the response body")
		return
	}
	var b bytes.Buffer
	errs := json.NewEncoder(&b).Encode(u)
	if errs != nil {
		log.Fatal("something went wrong deconding the response body")
		return
	}
	urls := "http://task1:2100/api/task1/id=21?game=soccer&tornament=premier"
	parsed, err := url.Parse(urls)
	if err != nil {
		log.Fatal("something went parsing urls")
		return
	}
	fmt.Println(parsed.Scheme, parsed.Path, parsed.Host, parsed.Query(), parsed.Port(), parsed.RequestURI())

}
