package main

import "fmt"

type Animal struct {
	Type string
}

func (a *Animal) Eat(str, food string) {
	fmt.Printf("%s Eats %s \n", str, food)
}
func (a *Animal) Lives(str string, d int) {
	fmt.Printf("%s can live upto  %d \n", str, d)
}

type Man struct {
	Name string
	*Animal
}

func NewMan(name, kind string) *Man {
	return &Man{
		Name:   name,
		Animal: &Animal{Type: kind},
	}
}
func (a *Man) Speak(str, lang string) {
	fmt.Printf("%s can speak %s \n", str, lang)
}

func main() {
	bob := NewMan("bob", "man")
	bob.Eat(bob.Name, "meat")
}
