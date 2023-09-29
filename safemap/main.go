package main

import (
	"fmt"
	"sync"
)

type SafeMap[K comparable, V any] struct {
	locker sync.RWMutex
	data   map[K]V
}

func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}
func (s *SafeMap[K, V]) Insert(key K, value V) {
	s.locker.Lock()
	defer s.locker.Unlock()
	s.data[key] = value
}
func (s *SafeMap[K, V]) Get(key K) (V, error) {
	s.locker.RLock()
	defer s.locker.RUnlock()
	value, ok := s.data[key]
	if !ok {
		return value, fmt.Errorf("%v not found", key)
	}
	return value, nil
}
func (s *SafeMap[K, V]) Update(key K, value V) error {
	s.locker.Lock()
	defer s.locker.Unlock()
	_, ok := s.data[key]
	if !ok {
		return fmt.Errorf("%v not found", key)
	}
	s.data[key] = value
	return nil
}

func (s *SafeMap[K, V]) Delete(key K) error {
	s.locker.Lock()
	defer s.locker.Unlock()
	_, ok := s.data[key]
	if !ok {
		return fmt.Errorf("%v not found", key)
	}
	delete(s.data, key)
	return nil
}
func (s *SafeMap[K, V]) HasKey(key K) bool {
	s.locker.Lock()
	defer s.locker.Unlock()
	_, ok := s.data[key]
	return ok
}
func main() {
	mapa := NewSafeMap[int, string]()
	for i := 1; i <= 10; i++ {
		go func(i int) {
			mapa.Insert(i, fmt.Sprintf(" goroutine %d working", i))
		}(i) 
	}
	for j := 1; j <= 10; j++ {
		v, _ := mapa.Get(j)
		fmt.Println(v)
	}

}
