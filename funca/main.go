package main

import "fmt"

type FuncInt func([]int) int

type store struct {
	db []int
}

func NewStore(d []int) *store {
	return &store{db: d}
}
func (s *store) Handler(f func([]int) int) {
	fmt.Println(f(s.db))
}
func getTotal(d []int) (sum int) {
	for _, k := range d {
		sum += k
	}
	return sum
}
func getMean(d []int) int {
	return getTotal(d) / len(d)
}
func getMax(d []int) int {
	max := d[0]
	for _, k := range d {
		if k > max {
			max = k
		}
	}
	return max
}
func getMin(d []int) int {
	min := d[0]
	for _, k := range d {
		if k < min {
			min = k
		}
	}
	return min
}
func main() {
	slice := []int{2, 4, 3, 5, 9, 7, 8, 6}
	store := NewStore(slice)
	store.Handler(getMax)
	store.Handler(getMean)
	store.Handler(getTotal)
	store.Handler(getMin)
}

// type FuncString func(string) string

// type Person struct {
// 	Name  string
// 	Funca FuncString
// }

// func NewPerson(name string) *Person {
// 	return &Person{Name: name}
// }
// func (p *Person) handler() {
// 	fmt.Println(p.Funca(p.Name))
// }
// func Write(name string) string {
// 	return fmt.Sprintf("%s can Write", name)
// }
// func Painter(name string) string {
// 	return fmt.Sprintf("%s can Paint", name)
// }
// func Walker(name string) string {
// 	return fmt.Sprintf("%s can Walk", name)
// }
// func main() {
// 	tony := NewPerson("Tony")
// 	tony.Funca = Write
// 	tony.handler()
// 	tony.Funca = Painter
// 	tony.handler()
// 	tony.Funca = Walker
// 	tony.handler()
// }
