package main

import "fmt"

type Team struct {
	Members []string
}
type Squard struct {
	Teams []string
}

func main() {
	var m1 map[string]int //declare a map
	// initilaizing a map
	m2 := make(map[string]int, 5) // permits one to create a mpa with length
	m3 := map[string]int{}        //permits one to create a map with values

	res1 := make(map[string][]int)
	res1["vamos"] = []int{1, 2, 3}
	fmt.Println("a map with slice of integers", res1)
	res3 := make(map[string][3]int)
	res3["vamos"] = [3]int{1, 2, 3}
	fmt.Println("a map with array of integers", res3)
	res4 := make(map[[2]int]string)
	res4[[2]int{1, 2}] = "escuela"
	fmt.Println("a map with array of integers", res4)

	fmt.Println(m1, m2, m3)
	m4 := map[string]int{"tony": 30, "jenny": 27, "mark": 25}
	fmt.Println(m4)
	m4["gutierez"] = 25
	fmt.Println(m4)
	delete(m4, "mark")
	fmt.Println(m4)
	// merging maps
	first := map[string]int{"a": 1, "b": 2, "c": 3}
	second := map[string]int{"a": 5, "d": 6, "e": 7, "c": 8}
	for k, v := range second {
		first[k] = v
	}
	fmt.Println(first)

	var users map[string]int
	users2 := map[string]int{"tony": 30, "jenny": 27, "mark": 25}
	users = users2
	fmt.Println(users)
	delete(users, "mark")
	fmt.Println(users)
	fmt.Println(users2)
	users3 := map[string]int{"tony": 30, "jenny": 27, "mark": 25}

	users4 := map[string]int{}
	for k, v := range users3 {
		users4[k] = v
	}
	delete(users3, "mark")
	fmt.Println(users3)
	fmt.Println(users4)
	////////////////////////multibranch map/////////////
	map1 := map[int]map[string]interface{}{}
	map1[1] = map[string]interface{}{
		"S1": Squard{[]string{"T1", "T2"}},
		"S2": Squard{[]string{"T3", "T4"}},
	}
	map1[2] = map[string]interface{}{
		"T1": Team{[]string{"Tony", "Jeny", "Fred"}},
		"T2": Team{[]string{"Gutierez", "Alex", "Nat"}},
		"T3": Team{[]string{"Santitago", "Mary", "Pablo"}},
	}
	fmt.Println(map1)

}
