package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

// Dining philosophers
const hunger = 3

// variables - philosophers
var (
	philosophers  = []string{"tony", "jenny", "marry", "gutierez", "santiago"}
	sleepTime     = 1 * time.Second
	eatTime       = 2 * time.Second
	thinkTime     = 1 * time.Second
	orderFinished []string
	wg            sync.WaitGroup
	orderMutex    sync.Mutex
)

func diningProblem(phil string, forkLeft, forkRight *sync.Mutex) {
	defer wg.Done()
	fmt.Println(phil, " is seated at the table")
	for i := hunger; i > 0; i-- {
		fmt.Println(phil, " is hungry")
		time.Sleep(sleepTime)
		// lock both forks
		forkLeft.Lock()
		fmt.Printf("\t%s picked up the fork to his  \n", phil)
		forkRight.Lock()
		fmt.Printf("\t%s picked up the fork to his right\n", phil)

		fmt.Println(phil, "is eating using both forks")
		time.Sleep(eatTime)

		fmt.Println(phil, "is thinking")
		time.Sleep(thinkTime)

		// unluck both forks
		forkLeft.Unlock()
		forkRight.Unlock()
		fmt.Printf("\t%s is done with the right fork\n", phil)
		fmt.Printf("\t%s is done with the left fork\n", phil)
		time.Sleep(sleepTime)
	}
	orderMutex.Lock()
	orderFinished = append(orderFinished, phil)
	orderMutex.Unlock()
	fmt.Println(phil, "is done eating")
	time.Sleep(sleepTime)
	fmt.Println(phil, "ihas left the table")

}
func main() {
	log.Println("Dining philosophers problem")
	// spawn go routines
	forkLeft := &sync.Mutex{}
	wg.Add(len(philosophers))
	for _, phil := range philosophers {
		forkRight := &sync.Mutex{}
		go diningProblem(phil, forkLeft, forkRight)
		forkLeft = forkRight
	}
	wg.Wait()
	fmt.Printf("Order finished %s \n", strings.Join(orderFinished, ", "))
}
