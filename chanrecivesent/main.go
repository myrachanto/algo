package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("barber challage")
	ping, pong := make(chan string), make(chan string)
	go shout(ping, pong)
	fmt.Println("type something and and q to quit")
	for {
		fmt.Print("-> ")
		var userInput string
		_, _ = fmt.Scanln(&userInput)
		if userInput == strings.ToLower("q") {
			break
		}
		ping <- userInput
		response := <-pong
		fmt.Println("Response: ", response)
	}
	fmt.Println("All done! closing the chanels")
	close(ping)
	close(pong)
}
func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping
		pong <- fmt.Sprintf("%s!", strings.ToUpper(s))
	}
}
