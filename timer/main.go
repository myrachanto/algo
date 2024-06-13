package main

import (
	"fmt"
	"time"
)

func main() {now := time.Now()

	// Calculate the time for yesterday
	yesterday := now.AddDate(0, 0, -1)
	lastMidnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	fmt.Println(lastMidnight, "-----", now)
}
