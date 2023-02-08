package main

import (
	"fmt"
	"sync"
)

type Trip struct {
	Driver string
	Hotel  string
	Rate   float64
}
type Driver struct {
	Name    string
	Rate    float64
	Counter int
}
type Store struct {
	Db map[string][]Trip
	Mu sync.Mutex
}

func NewStore() *Store {
	return &Store{
		Db: make(map[string][]Trip),
		Mu: sync.Mutex{},
	}
}

func main() {
	trips := []Trip{
		{"tony", "casa branka", 4},
		{"john", "brianka", 3},
		{"mary", "Lenatia", 5},
		{"mary", "cas branka", 4},
		{"tony", "brianka", 5},
		{"tony", "cas branka", 4},
	}
	db := NewStore()
	db.DriverAllrates(trips)
	fmt.Println("mapping ------------", db.Db)
	drivers := db.Driverrates()
	fmt.Println("drivers and their rates ------------", drivers)
	besdriver := findBestDriver(drivers)
	fmt.Printf("The best driver is %s with a rate of %.1f with %d number of trips \n", besdriver.Name, besdriver.Rate, besdriver.Counter)
	fmt.Println("////////////////////////////////////")
	results := DriverAllrates(trips)
	fmt.Println("mapping ------------", results)
	drivers2 := Driverrates(results)
	fmt.Println("drivers and their rates ------------", drivers2)
	besdriver2 := findBestDriver(drivers2)

	// bestdriver := findBestDriver(trips)
	fmt.Printf("The best driver is %s with a rate of %.1f with %d number of trips \n", besdriver2.Name, besdriver.Rate, besdriver.Counter)
	// averagerate := findAvaragerating(trips)
	// fmt.Printf("Average rate is %d", averagerate)
}

func DriverAllrates(data []Trip) map[string][]Trip {
	mapa := make(map[string][]Trip)
	for _, trip := range data {
		mapa[trip.Driver] = append(mapa[trip.Driver], trip)
	}
	return mapa
}

func (s *Store) DriverAllrates(data []Trip) {
	for _, trip := range data {
		s.Mu.Lock()
		s.Db[trip.Driver] = append(s.Db[trip.Driver], trip)
		s.Mu.Unlock()
	}
}
func (s *Store) Driverrates() []Driver {
	results := []Driver{}
	for name, trip := range s.Db {
		res := Driver{}
		for _, t := range trip {
			res.Name = name
			res.Counter++
			res.Rate += t.Rate
		}
		results = append(results, res)
	}
	return results
}

func Driverrates(data map[string][]Trip) []Driver {
	results := []Driver{}
	for name, trip := range data {
		res := Driver{}
		for _, t := range trip {
			res.Name = name
			res.Counter++
			res.Rate += t.Rate
		}
		results = append(results, res)
	}
	return results
}
func findBestDriver(data []Driver) *Driver {
	maxrate := data[0].Rate / float64(data[0].Counter)
	bestdriver := &Driver{}
	for _, driver := range data {
		res := float64(driver.Rate) / float64(driver.Counter)
		if res > maxrate {
			maxrate = res
			bestdriver.Name = driver.Name
			bestdriver.Rate = res
			bestdriver.Counter = driver.Counter
		}
	}

	return bestdriver
}

// func findBestDriverSlice(data []Driver) (sl []float64){
// 	maxrate := data[0].Rate
// 	bestdirver := Driver{}
// 	for _, driver := range data {
// 		res := float64(driver.Rate) / float64(driver.Counter)
// 		if res > maxrate {
// 			maxrate = res
// 			bestdirver = driver
// 		}
// 	}
// 	return &bestdirver, maxrate
// }

// func getmax(nums []float64) float64 {
// 	max := nums[0]
// 	for _, g := range nums {
// 		if max < g {
// 			max = g
// 		}
// 	}
// 	return max
// }

// func findAvaragerating(data []Trip) int {
// 	rates := 0
// 	for _, driver := range data {
// 		rates += driver.Rate

// 	}
// 	return rates / len(data)
// }
