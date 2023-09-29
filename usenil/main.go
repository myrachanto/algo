package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) String() string {
	return fmt.Sprintf("My name is %s \n", p.Name)
}

func main() {
	var p *Person
	var s []string

	// fmt.Println(p == s) // will not compile

	var m map[string]int
	m["felicity"] = 0 //will panic
	fmt.Println(m)
	s = append(s, "John Diego")
	fmt.Println(s)

	p = &Person{"Oliver Queen"}
	fmt.Println(p)

	var str fmt.Stringer = p
	fmt.Println(str == nil)

}
