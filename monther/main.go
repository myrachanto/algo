package main

import (
	"fmt"
	"time"
)

var holidays = []time.Time{
	time.Date(2023, time.May, 10, 23, 0, 0, 0, time.UTC),
	time.Date(2023, time.May, 15, 23, 0, 0, 0, time.UTC),
	time.Date(2023, time.May, 20, 23, 0, 0, 0, time.UTC),
}

func main() {
	end := getnextpayday(time.Now(), 30)
	fmt.Println("----------------------")
	fmt.Println(end)
}

func getnextpayday(start time.Time, days int) time.Time {
	var j int
	for i := 0; i < days; {
		if start.Weekday() != time.Saturday && start.Weekday() != time.Sunday {
			for _, v := range holidays {
				ok := checkholiday(start, v)
				if ok {
					j++
				}
			}
			i++
		}
		start = start.AddDate(0, 0, 1)
	}
	if j > 0 {
		fmt.Println("----------------------", j)
		start = start.AddDate(0, 0, j)
	}
	return start
}
func checkholiday(day time.Time, holi time.Time) bool {
	return day.YearDay() == holi.YearDay()
}
