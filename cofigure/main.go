package main

import "fmt"

type Server struct {
	Opts
}
type Opts struct {
	Id      int
	Name    string
	MaxConn int
	Timeout int
	TLS     bool
}

func New(opts ...OptFunc) *Server {
	options := DefaultOpts()
	for _, fn := range opts {
		fn(options)
	}
	return &Server{*options}
}

type OptFunc func(*Opts)

func DefaultOpts() *Opts {
	return &Opts{
		Name:    "Default",
		Id:      1902,
		MaxConn: 10,
		Timeout: 1,
		TLS:     false,
	}
}

func WithTLS(o *Opts) {
	o.TLS = true
}

func WithName(name string) OptFunc {
	return func(o *Opts) {
		o.Name = name
	}
}

func main() {
	s := New(WithTLS, WithName("Golang server"))
	fmt.Printf("%#v \n", s)
}
