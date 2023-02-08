package main

import (
	"fmt"
)

type (
	Reader interface {
		Read()
	}
	Writer interface {
		Write(buf []byte) (int, error)
	}
	String interface {
		String(string) string
	}
	Person interface {
		Reader
		Writer
		String
	}
	person struct {
		name string
		age  int
		Person
		// address func(String)
	}
)

func (p *person) String(str string) string {
	return fmt.Sprintf("%s is %d and as such is qualified for this %s position", p.name, p.age, str)
}
func (p *person) Read() {
	n, _ := fmt.Println("welcome This is in the reader interface")
	fmt.Println(n)
}
func NewPerson(name string, age int) Person {
	return &person{name: name, age: age}
}
func main() {
	// var e interface{}
	// e = 7
	// fmt.Printf("------------ value is %v of type %T \n", e, e)
	// str1, ok := e.(string)
	// if ok {
	// 	fmt.Printf("------------ value is %v of type %T \n", str1, str1)
	// }
	// e = "my name"
	// fmt.Printf("------------ value is %v of type %T\n", e, e)
	// str, ok := e.(string)
	// if ok {
	// 	fmt.Printf("------------ value is %v of type %T \n", str, str)
	// }
	///////////////////////////////////////////////////////////
	tony := NewPerson("tony", 30)
	fmt.Println(tony.String("CTO"))
	tony.Read()
}
