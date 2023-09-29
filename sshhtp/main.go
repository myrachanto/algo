package main

import (
	"io"
	"log"
	"sync"

	"github.com/gliderlabs/ssh"
)

type Tunnel struct {
	Writer io.Writer
	DoneCh chan struct{}
}
type Store struct {
	Db     map[int]Tunnel
	Locker sync.Mutex
}

func New() *Store {
	return &Store{
		Db: make(map[int]Tunnel),
	}
}

func init() {
	log.SetPrefix("SSH: =>")
}

func main() {
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, "Hello world\n")
	})
	log.Println("-----------------")
	port := ":2222"
	log.Fatal(ssh.ListenAndServe(port, nil))
	log.Println("Server started at port", port)
}
