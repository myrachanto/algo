package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dollars float32

type Store struct {
	db     map[string]dollars
	locker sync.Mutex
}

func NewDatabase(m map[string]dollars) *Store {
	return &Store{
		db:     m,
		locker: sync.Mutex{},
	}
}

func (d dollars) String() string {
	return fmt.Sprintf("KSH%.2f", d)
}
func (s *Store) Create(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")
	s.locker.Lock()
	if _, ok := s.db[item]; ok {
		msg := fmt.Sprintf("duplicate item: %q ", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("Invalid price: %q ", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	s.db[item] = dollars(p)
	s.locker.Unlock()
	fmt.Fprintf(w, "Item added succesfully")
}
func (s *Store) GetOne(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	s.locker.Lock()
	if _, ok := s.db[item]; !ok {
		msg := fmt.Sprintf("item not found: %q ", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	s.locker.Unlock()
	fmt.Fprintf(w, "Item found %s with %v", item, s.db[item])
}
func (s *Store) GetAll(w http.ResponseWriter, r *http.Request) {
	s.locker.Lock()
	for item, price := range s.db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
	s.locker.Unlock()
}
func (s *Store) Update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")
	s.locker.Lock()
	if _, ok := s.db[item]; !ok {
		msg := fmt.Sprintf("item not found: %q ", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("Invalid price: %q ", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	s.db[item] = dollars(p)
	s.locker.Unlock()
	fmt.Fprintf(w, "Item Updated succesfully")
}
func (s *Store) Delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	s.locker.Lock()
	if _, ok := s.db[item]; !ok {
		msg := fmt.Sprintf("item not found: %q ", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	delete(s.db, item)
	s.locker.Unlock()
	fmt.Fprintf(w, "%v deleted successfully", item)
}
func main() {
	d := map[string]dollars{"shoes": 500, "socks": 100}
	db := NewDatabase(d)
	http.HandleFunc("/create", db.Create)
	http.HandleFunc("/getone", db.GetOne)
	http.HandleFunc("/getall", db.GetAll)
	http.HandleFunc("/update", db.Update)
	http.HandleFunc("/delete", db.Delete)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
