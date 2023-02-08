package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	money          = 100
	lock           = sync.Mutex{}
	wg             = sync.WaitGroup{}
	moneyDeposited = sync.NewCond(&lock)
	counter1       = 0
	counter2       = 0
)

// conditional sync
func stingy() {
	fmt.Println("stingy started")
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money += 10
		counter1++
		fmt.Println("Stingy the balance is ", money)
		moneyDeposited.Signal()
		lock.Unlock()
		time.Sleep(time.Millisecond * 1)
	}
	wg.Done()
	fmt.Println("stingy done")
}
func spendy() {
	fmt.Println("Spendy started")
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		for money-20 < 0 {
			moneyDeposited.Wait()
		}
		money -= 20
		counter2++
		fmt.Println("Spendy the balance is ", money)
		lock.Unlock()
		time.Sleep(time.Millisecond * 1)
	}
	wg.Done()
	fmt.Println("Spendy done")
}
func main() {
	wg.Add(1)
	go stingy()
	go spendy()  
	wg.Wait()
	fmt.Println("Balance ", money, counter1, counter2)
}
