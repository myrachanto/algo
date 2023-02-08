package main

import "fmt"

type item struct{
	name string
	product int
}
func main(){
	// Task
	// find the most popular product 

	mydata := map[string]item{
		"Mike": item{"nike", 1},
		"Sara": item{"Ribox", 1},
		"Anna": item{"calvin", 1},
		"Peter": item{"nike", 1},
		"Luke": item{"nike", 1},
		"Jenn": item{"Ribox", 1},
	}
	getpops := getmap(mydata)
	res := getkey(getpops)
	sorted := sort(res)
	mostpopular := sorted[len(sorted) -1]
	popularshoe := getpopularshoe(mostpopular, getpops)
	fmt.Println("The msot Popular shoe is ", popularshoe)

}
func getpopularshoe(mostpopular int, data map[string]int) string{
	popularshoe := ""
	for m,n := range data{
		if n == mostpopular {
			popularshoe = m
		}
	} 
	return popularshoe
}
func getmap(data  map[string]item)map[string]int{
	res := map[string]int{}
	for _,y := range data{
		res[y.name]++
	}
	return res
} 
func getkey(data map[string]int) []int{
	res := []int{}
	for _,f := range data{
		res = append(res, f)
	}
	return res
}
func sort(data []int) []int{
	for i:=0; i<len(data); i++{
		for j:= i + 1; j< len(data); j++{
			if data[i] > data[j]{
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}