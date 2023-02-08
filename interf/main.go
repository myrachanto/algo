package main

import "fmt"

type User struct {
	name, address string
	age           uint8
}
type UserInterface interface {
	SetName(name string)
	GetName() string
	SetAge(age uint8)
}

func NewUser(name string, age uint8) UserInterface {
	return &User{name: name, age: 32}
}

func (u *User) SetName(name string) {
	u.name = name
}
func (u User) GetName() string {
	return u.name
}
func (u *User) SetAge(age uint8) {
	u.age = age
}
func (u User) String() string {
	return fmt.Sprintf("%s is %d years old", u.name, u.age)
}
func main() {
	// var u1 UserInerface = &User{}
	u1 := NewUser("chantos", 27)
	u1.SetName("anthony")
	// fmt.Println(u1.GetName())
	fmt.Printf("%T of %v is %#v \n", u1, u1, u1)
	fmt.Println(u1)
}
