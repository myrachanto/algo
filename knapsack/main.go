package main

import "fmt"

type (
	Store struct {
		Db []Itm
	}
	Itm struct {
		Profit int
		Weight int
		PW     float64
	}
)

func NewInstance(data []Itm) *Store {
	return &Store{Db: data}
}

func main() {
	sl := []Itm{
		{10, 2, 10 / 2},
		{5, 3, 5 / 3},
		{15, 5, 15 / 5},
		{7, 7, 1},
		{6, 1, 6 / 1},
		{18, 4, 18 / 4},
		{4, 1, 3 / 1}}
	store := NewInstance(sl)
	maxweight := 15.0
	pwslice := store.convertoslices()
	sorter(pwslice)
	fmt.Println(pwslice)
	psortedpw := getbestresults(maxweight, pwslice)
	fmt.Println(psortedpw)
	for k, item := range store.getItems(psortedpw) {
		fmt.Printf("Item %d of weight %d by profit %d\n", k, item.Weight, item.Profit)
	}

}
func (s *Store) convertoslices() []float64 {
	pwslice := []float64{}
	for _, v := range s.Db {
		pwslice = append(pwslice, v.PW)
	}
	return pwslice
}
func getbestresults(maxweight float64, data []float64) []float64 {
	psortedpw := []float64{}
	for _, v := range data {
		if sum(psortedpw) >= maxweight {
			break
		}
		psortedpw = append(psortedpw, v)
	}
	return psortedpw
}

func (s *Store) getItems(data []float64) []Itm {
	res := Itm{}
	results := []Itm{}
	for _, d := range data {
		for _, k := range s.Db {
			if d == k.PW {
				res = k
			}
		}
		results = append(results, res)
	}
	return results
}
func sum(data []float64) float64 {
	sum := 0.0
	for _, g := range data {
		sum += g
	}
	return sum
}
func sorter(data []float64) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] < data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
}
