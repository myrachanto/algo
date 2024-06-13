package main

import "sync"

type Something interface {
	Read(b []byte) (int, error)
	Write(b []byte) (int, error)
}
type something struct {
	name            string
	sublocker       *sync.RWMutex
	subscribers     []chan map[string]message
	foolowinglocker *sync.Mutex
	following       []string
	Something
	age int
}
type message struct {
	name     string
	location string
	likes    int
}

// type vamo
