package main

import (
	"fmt"
	"strings"
	"time"
)

type Minutes int
type User struct {
	Name, Username, Email, Phone string
	Age, Id                      uint8
	Salary                       float32
	Address                      Address
}
type Address struct {
	Location, Country string
}

func main() {
	location, _ := time.LoadLocation("America/New_york")
	fmt.Println(location, time.Now().In(location))
	year, month, day := time.Now().Date()
	fmt.Println(strings.Repeat("/", 50))
	fmt.Printf("%7v| %9v|%8v|\n", "Year", "Month", "Day")
	fmt.Printf("%7v|  %7v| %7v|\n", year, month, day)
	fmt.Println(strings.Repeat("/", 50))
	fmt.Println(strings.Repeat("*", 50))
	fmt.Printf("%-7v| %-9v|%-8v|\n", "Year", "Month", "Day")
	fmt.Printf("%-7v|  %-7v| %-7v|\n", year, month, day)
	fmt.Println(strings.Repeat("/", 50))
	munites := Minutes(45)
	fmt.Println(munites)
	dev1 := User{Name: "tony", Username: "chantos", Email: "myrachanto@gmail.com", Phone: "25435677", Age: 31, Id: 1, Address: Address{Country: "kenya", Location: "Nairobi"}}
	fmt.Println(dev1)

}
