package main

import "fmt"

type FuncString func(string) string

type Person struct {
	Name  string
	Funca FuncString
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}
func (p *Person) handler() {
	fmt.Println(p.Funca(p.Name))
}
func Write(name string) string {
	return fmt.Sprintf("%s can Write", name)
}
func Painter(name string) string {
	return fmt.Sprintf("%s can Paint", name)
}
func Walker(name string) string {
	return fmt.Sprintf("%s can Walk", name)
}
func main() {
	tony := NewPerson("Tony")
	tony.Funca = Write
	tony.handler()
	tony.Funca = Painter
	tony.handler()
	tony.Funca = Walker
	tony.handler()
}
