package main

import (
	"fmt"
	"sync"
)

func main() {
	syncmap := sync.Map{}
	// Storing data int sync map
	syncmap.Store(0, "tony")
	syncmap.Store(1, "Jenny")
	// retriving data from sync map
	value, ok := syncmap.Load(0)
	fmt.Println(value, ok)
	// deleting items from sync map
	syncmap.Delete(0)
	// fmt.Println(syncmap)
	syncmap.Store(0, "tony")
	//deleting and loading the deleted item
	value, loaded := syncmap.LoadAndDelete(1)
	fmt.Println(value, loaded)
	syncmap.Store(1, "Jenny")
	// if not exist add it to the map
	value, load := syncmap.LoadOrStore(2, "calorina")
	fmt.Println(value, load)
	var wg sync.WaitGroup
	syncmap2 := sync.Map{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			syncmap2.Store(i, fmt.Sprintf("working %d", i)) 
			wg.Done()
		}(i)
	}
	wg.Wait()
	syncmap2.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
