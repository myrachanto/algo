package main

import "sync"

func main() {
	total, max := 10, 3
	var wg sync.WaitGroup
	for i := 0; i < total; i += max {
		limit := max
		if i+max > total {
			limit = total - i
		}
		wg.Add(limit)
		for j := 0; j < limit; j++{

		}

	}
}
