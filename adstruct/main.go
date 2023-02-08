package main

import "fmt"

type User struct {
	Username    string
	Connections map[string]int
}

func main() {
	p0 := &User{"Tony", map[string]int{"fred": 2345, "Jenny": 3532}}
	fmt.Println(p0)
	p1 := []User{
		{"Tony", map[string]int{"fred": 2345, "Jenny": 3532}},
		{"vamos", map[string]int{"bairal": 2345, "camino": 3532}},
	}
	for _, i := range p1 {
		fmt.Println(i.Username)
		for _, v := range i.Connections {
			fmt.Printf("%10v \n", v)

		}
	}
	
}
