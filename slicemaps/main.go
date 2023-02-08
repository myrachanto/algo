package main

import (
	"fmt"
)

func main() {
	ls := []int{4, 4, 5}
	ls = append(ls, 2)
	fmt.Println(ls)
	ls1 := [5]int{4, 4, 5}
	fmt.Println(ls1)
	fmt.Println(ls, len(ls), cap(ls))
	ls = append(ls, 3, 5, 6)
	fmt.Println(ls, len(ls), cap(ls))
	stsi := []string{"hoy", "voy"}
	fmt.Println(stsi, len(stsi), cap(stsi))
	stsi = append(stsi, "samo", "dc")
	fmt.Println(stsi, len(stsi), cap(stsi))
	sub := stsi[0:2]
	sub = append(sub, "vsmos", "scd", "dcs")
	fmt.Println(sub, len(sub), cap(sub))
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sl2 := []int{}
	sl2s := []int{}
	copy(sl,sl2s)//doesnt work!!!!!
	fmt.Println("coping slices ---------------", sl2s)
	for _, j := range sl {
		if j == 5 {
			continue
		}
		sl2 = append(sl2, j)
	}
	fmt.Println(sl2)
	var i int = 0
l1:
	i++
	if i <= 10 {
		fmt.Println("vamos ---", i)
		goto l1
	} else {
		fmt.Println("done with the l1")
	}
	slice1 := []string{"hoy", "te", "amo"}
	slice2 := []int{1, 2, 3}
	fmt.Println(mapper(slice1, slice2))
	str := "Anthony"
	fmt.Println("My Name start with --->>>>", str[1])
	fmt.Println("My Name start with --->>>>", str[:3])
	str += "s"
	fmt.Println(str)
	a := 2 
	fmt.Printf("a: %8T %[1]v\n", a)
	v := make([]int, 4, 4)
	v = append(v, 5) //this will append from index 5 [00005]
	fmt.Println(v)

	d3 := []string{"tony", "mark", "june", "jane", "jenny"}
	d4 := []string{}
	for _, j := range d3 {
		// time.Sleep(1 * time.Second)
		d4 = append(d4, j)
	}
	fmt.Println(d3)
	fmt.Println(d4)

}
func mapper(slice1 []string, slice2 []int) map[string]int {
	if len(slice1) != len(slice2) {
		return nil
	}
	n := map[string]int{}
	for i, v := range slice1 {
		n[v] = slice2[i]
	}
	return n
}
