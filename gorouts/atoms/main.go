package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg = &sync.WaitGroup{}

type calculator struct {
	res atomic.Value
}

func newCalculator() calculator {
	c := calculator{}
	c.res.Store(float64(0))
	return c
}
func (c *calculator) Add(n float64) {
	c.res.Store(c.results() + n)
}
func (c *calculator) Sub(n float64) {
	c.res.Store(c.results() - n)

}
func (c *calculator) Mul(n float64) {
	c.res.Store(c.results() * n)

}
func (c *calculator) Div(n float64) {
	if n == 0 {
		panic("cannot divide by zero")
	}
	c.res.Store(c.results() / n)

}
func (c *calculator) results() float64 {
	r, ok := c.res.Load().(float64)
	if !ok {
		panic("operating with the wrong type")
	}
	return r

}

func main() {
	//doesnt work
	c := newCalculator()
	wg.Add(5)
	go func() {
		defer wg.Done()
		c.Add(7)
	}()
	go func() {
		defer wg.Done()
		c.Mul(2)
	}()
	go func() {
		defer wg.Done()
		c.Sub(7)
	}()
	go func() {
		defer wg.Done()
		c.Div(4)
	}()
	go func() {
		defer wg.Done()
		c.Add(10)
	}()
	wg.Wait()
	fmt.Println("--------------", c.results())
}
