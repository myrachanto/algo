package main

import (
	"fmt"
	"sync"
)

type Task struct {
	Id int
}

func worker(wg *sync.WaitGroup, work chan Task, goid int) {
	defer wg.Done()
	// time.Sleep(time.Duration(goid) * time.Millisecond)
	for task := range work {
		fmt.Printf("work done for task with id %d done by go routine %d\n", task.Id, goid)
	}
}
func main() {
	var wg sync.WaitGroup
	work := make(chan Task, 10)
	workerpool := 3
	//  start workers
	for i := 1; i <= workerpool; i++ {
		wg.Add(1)
		go worker(&wg, work, i)
	}
	for i := 1; i <= 10; i++ {
		work <- Task{Id: i}
	}
	close(work)
	wg.Wait()

}
