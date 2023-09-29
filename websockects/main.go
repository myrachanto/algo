package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	locker sync.Mutex
	conns  map[*websocket.Conn]bool
}

func New() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}
func (s *Server) HandleWs(ws *websocket.Conn) {
	fmt.Printf("Incomming connection from a client - %v \n", ws)
	s.locker.Lock()
	s.conns[ws] = true
	s.locker.Unlock()
	s.readloops(ws)
}
func (s *Server) Orders(ws *websocket.Conn) {
	fmt.Println(" orders as they are comming")
	for {
		payload := fmt.Sprintf("Order flowing %d\n", time.Now().Unix())
		ws.Write([]byte(payload))
		time.Sleep(1 * time.Second)
	}

}
func (s *Server) readloops(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error occured %v\n", err)
			continue
		}
		msg := buf[:n]
		fmt.Println(string(msg))
		s.broadcast(msg)
		ws.Write([]byte("muchas gracias"))
	}
}
func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("error occred while trying to broadcast", err)
			}
		}(ws)
	}
}
func main() {
	fmt.Println("server started......")
	s := New()
	http.Handle("/ws", websocket.Handler(s.HandleWs))
	http.Handle("/orders", websocket.Handler(s.Orders))
	http.ListenAndServe(":7000", nil)
}
