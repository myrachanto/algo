package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type Server struct {
	dblocker    sync.Mutex
	db          map[int]User
	dbhit       int
	cachelocker sync.Mutex
	cache       map[int]User
}
type User struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func New() *Server {
	var db = make(map[int]User)
	var cache = make(map[int]User)
	user := User{}
	for i := 1; i <= 100; i++ {

		user.Id = i
		user.Name = fmt.Sprintf("username_%d", i)
		db[i] = user
	}
	return &Server{
		db:    db,
		cache: cache,
	}
}
func (s *Server) HandleGetUserById(w http.ResponseWriter, r *http.Request) {
	idstr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		panic("could not get the id")
	}
	user, ok := s.cache[id]
	if ok {
		json.NewEncoder(w).Encode(&user)
		return
	}
	user, ok = s.db[id]
	if !ok {
		panic("could not get the user")
	}
	s.dbhit++
	s.cachelocker.Lock()
	s.cache[id] = user
	s.cachelocker.Unlock()
	json.NewEncoder(w).Encode(&user)
}
func main() {

}
