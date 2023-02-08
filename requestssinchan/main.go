package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var (
	wg  sync.WaitGroup
	url = "https://jsonplaceholder.typicode.com/users"
)

type Data struct {
	Name     string      `json:"name,omitempty"`
	Username string      `json:"username,omitempty"`
	Email    string      `json:"email,omitempty"`
	Id       int         `json:"id,omitempty"`
	Address  interface{} `json:"address,omitempty"`
}

func requester(j, i int) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	data := []Data{}
	err1 := json.Unmarshal([]byte(body), &data)
	if err1 != nil {
		return
	}
	fmt.Println("request proceesed", j+1+i)
}
func main() {
	// requester()
	total, max := 10, 3
	for i := 0; i < total; i += max {
		limit := max
		if i+max > total {
			limit = total - i
		}
		wg.Add(limit)
		for j := 0; j < limit; j++ {
			go requester(j, i)
		}
		time.Sleep(1 * time.Second)
		fmt.Println("--------------------")
		wg.Wait()
	}
}
