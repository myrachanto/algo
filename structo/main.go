package main

import "fmt"

type Person struct {
	Age  int
	Role bool
}

func updaterole(p Person) {
	p.Role = true
}
func checkrole(p *Person) bool {
	return p.Role
}
func main() {
	p := &Person{}
	updaterole(*p)
	fmt.Println(checkrole(p))
}
