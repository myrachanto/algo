package main

import (
	"fmt"
	"sort"
	"sync"
)

type Person struct {
}

type Report struct {
	key   string
	value int
}

type D int
type Db struct {
	Store []Report
	Size  int
}

func main() {

	data := []int{1, 2, 3}
	sum := getsum(data)
	fmt.Println("get sum", sum)
	testar()
	// bs := make([]int, 0, 10)
	var bs []int
	bs = append(bs, data...)
	fmt.Printf("%v of capacity %d with the length of %d memory %p \n", bs, cap(bs), len(bs), bs)
	testslice(&bs)
	fmt.Println(bs)
	fmt.Printf("%v of capacity %d with the length of %d memory %p \n", bs, cap(bs), len(bs), bs)
}
func getsum(d []int) int {
	if len(d) == 0 {
		return 0
	}
	return d[0] + getsum(d[1:])
}
func testar() {
	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}
	a := [][]byte{}
	for k := range items {
		// a = append(a, v[:]) //doesnt work
		a = append(a, items[k][:])
	}
	fmt.Println("----------", a)
}
func testslice(b *[]int) {
	*b = append(*b, 4)
}
func otherstuff() {

	// var p *Person
	// fmt.Println(p == nil)
	// var p2 interface{}
	// fmt.Println(p2 == nil)

	// fmt.Println(p2 == p)
	var ds D
	fmt.Println(ds == 0)
	var d2 int
	fmt.Println(d2 == 0)

	// fmt.Println(ds == d2) // mismatched type!
	data := []int{1, 2, 3, 4, 5}
	sum := getsum(data)
	fmt.Println("get sum", sum)
	report := map[string]int{"tony": 1, "jenny": 5, "peter": 3, "carolina": 7}
	sli := []Report{}
	for k, v := range report {
		sli = append(sli, Report{k, v})
	}
	fmt.Println("slice", sli)
	sort.Slice(sli, func(i, j int) bool {
		return sli[i].value > sli[j].value
	})
	fmt.Println("sorted slice", sli)
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		i := i
		go func(i int) {
			fmt.Println(i, &i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	var p []int
	p = append(p, 5)
	fmt.Println(p)
	var g map[string]int
	// g["go"] = 1
	fmt.Println(g)
	fk := []int{1, 2, 3, 4, 5, 6, 7}
	fks := fk[:3]
	fkv := fks[:6] //will work
	fmt.Println(fkv)
	// fks = fk[:3:3]
	// fkv = fks[:6] //panics
	fmt.Println(fkv)
	db := new(Db)
	db.Store = append(db.Store, Report{"vamos", 7})
	fmt.Println(db)

	var report1 Report
	fmt.Printf("%T of %[1]v\n", report1)
	var report2 = new(Report)
	fmt.Printf("%T of %[1]v\n", report2)
}
