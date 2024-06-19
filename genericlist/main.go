package main

import "fmt"

type GenelicList[T comparable] struct {
	DB []T
}

func NewGenericList[T comparable]() *GenelicList[T] {
	return &GenelicList[T]{
		DB: []T{},
	}
}
func (g *GenelicList[T]) Insert(value T) {
	g.DB = append(g.DB, value)
}
func (g *GenelicList[T]) Remove1(value T) {
	for i, v := range g.DB {
		if v == value {
			lastIndex := len(g.DB) - 1
			g.DB[i] = g.DB[lastIndex]
			g.DB = g.DB[:lastIndex]
			return
		}
	}
}
func (g *GenelicList[T]) Remove(value T) {
	temp := []T{}
	for _, v := range g.DB {
		if v != value {
			temp = append(temp, v)
		}
	}
	g.DB = temp
}
func main() {
	db := NewGenericList[int]()
	db.Insert(1)
	db.Insert(2)
	db.Insert(3)
	db.Insert(4)
	db.Remove(3)
	fmt.Println(db)
}
