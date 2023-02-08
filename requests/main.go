package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var (
	wg     sync.WaitGroup
	locker sync.Locker
	url    = "https://jsonplaceholder.typicode.com/users"
)

type Data struct {
	Name     string      `json:"name,omitempty"`
	Username string      `json:"username,omitempty"`
	Email    string      `json:"email,omitempty"`
	Id       int         `json:"id,omitempty"`
	Address  interface{} `json:"address,omitempty"`
}
type Store struct {
	Data []Data
	// Quit chan chan error
}
type Db struct {
	Results      []Data
	Totalresults int
}

// func (s *Store) Close() error {
// 	ch := make(chan error)
// 	s.Quit <- ch
// 	return <-ch
// }

// request data
func (s *Store) requester(j,i int) {
	defer wg.Done()

	// ch := make(chan error)
	response, err := http.Get(url)
	if err != nil {
		// ch <- err
		// s.Quit <- ch
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// ch <- err
		// s.Quit <- ch
		return
	}
	data := []Data{}
	err1 := json.Unmarshal([]byte(body), &data)
	if err1 != nil {
		// ch <- err
		// s.Quit <- ch
		return
	}
	// locker.Lock()
	s.Data = append(s.Data, data...)
	// locker.Unlock()
	// fmt.Println("--------------------before stalling")
	// for _, v := range data {
	// 	// fmt.Println("--------------------before stalling", v)
	// 	s.Data <- v
	// }
	// fmt.Println("==============>>>>>>>>>>", data)
	fmt.Println("request proceesed", j+1+i)
}
func (s *Store) Compiler() (db Db) {
	// err <- s.Quit
	// if err
	fmt.Println("-------------------- Compiler")
	db.Results = append(db.Results, s.Data...)
	return db
}
func Writejsonfile(db []Data) error {

	file, err := json.MarshalIndent(&db, "", " ")
	if err != nil {
		fmt.Println("Unable to create json file")
		return err
	}

	return ioutil.WriteFile("db.json", file, 0644)
}
func main() {
	// requester()
	total, max := 10, 3
	store := &Store{}
	// res := make(chan Data, 100)
	for i := 0; i < total; i += max {
		limit := max
		if i+max > total {
			limit = total - i
		}
		wg.Add(limit)
		for j := 0; j < limit; j++ {
			go store.requester(j,i)
		}
		fmt.Println("--------------------")
		wg.Wait()
	}
	db := store.Data
	fmt.Println("Total results count of - ", len(db))
	err := Writejsonfile(db)
	if err != nil {
		fmt.Println(err)
	}
}
