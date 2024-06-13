package main

import "fmt"

type User struct {
	Name   string
	Career string
	Salary float32
}

func main() {
	// how to easily group careers
	users := []User{
		{Name: "Jenny", Career: "Golang Developer", Salary: 12000},
		{Name: "Anthony", Career: "Golang Developer", Salary: 8000},
		{Name: "Jack", Career: "DevOps", Salary: 10000},
		{Name: "Alexander", Career: "Marketing", Salary: 5000},
		{Name: "Carolina", Career: "Marketing", Salary: 7000},
	}
	groupdata := make(map[string][]User)
	for _, user := range users {
		groupdata[user.Career] = append(groupdata[user.Career], user)
	}
	fmt.Println(groupdata)
}
