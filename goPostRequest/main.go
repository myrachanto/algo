package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

func main() {
	posturl := "https://jsonplaceholder.typicode.com/posts"

	body := []byte(`{
		"title": "Post title",
		"body": "Post description",
		"userId": 1
	}`)

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	post := &Post{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}

	if res.StatusCode != http.StatusCreated {
		panic(res.Status)
	}

	fmt.Println("Id:", post.Id)
	fmt.Println("Title:", post.Title)
	fmt.Println("Body:", post.Body)
	fmt.Println("UserId:", post.UserId)
}

// r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
// if err != nil {
// 	panic(err)
// }

// r.Header.Add("Content-Type", "application/json")
// fmt.Println("----------------------------step 1")
// client := &http.Client{}
// res, err := client.Do(r)
// if err != nil {
// 	panic(err)
// }

// // fmt.Println("----------------------------step 2", string(res.Body))
// defer res.Body.Close()
// body, err = ioutil.ReadAll(res.Body)
// if err != nil {
// 	panic(err)
// }

// fmt.Println("----------------------------step 2", string(body))

// post := &Post{}
// fmt.Println("----------------------------step 3")
// derr := json.Unmarshal([]byte(body), &post)
// if derr != nil {
// 	panic(derr)
// }

// if res.StatusCode != http.StatusCreated {
// 	panic(res.Status)
