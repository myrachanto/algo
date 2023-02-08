package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Ball struct{ hits int }

func main() {
	requestChan := make(chan chan error)

	// starting the service gourountine
	go goroutine(requestChan)
	// Send 5 request and collect errors
	for i := 0; i < 5; i++ {
		fmt.Println("Request ", i+1)
		//  Make channel that will be used to communitacate error back to main
		errorchan := make(chan error)
		requestChan <- errorchan
		err := <-errorchan
		fmt.Printf("Eror received : %v\n", err)
		time.Sleep(1 * time.Second)
		if err != nil {
			break
		}

	}

}
func goroutine(requestChan <-chan chan error) {
	for g := range requestChan {
		fmt.Println("Got request from Requestchan error")
		g <- SomeOperation()
	}
	// for {
	// 	select {
	// 	case errChan := <-requestChan:
	// 		fmt.Println("Got request from Requestchan error")
	// 		errChan <- SomeOperation()
	// 	}
	// }
}
func SomeOperation() error {
	if rand.Intn(10) < 8 {
		return nil
	}
	return errors.New("error in SomeOperation")
}
func pingpong() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball)
	time.Sleep(1 * time.Second)
	<-table

}
func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
