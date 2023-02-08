package main

import (
	"fmt"
)

func main() {
	// map := make(map[int]int)
	// var mapa map[int]int = map[int]int{}
	myarray := map[int]int{0: 5, 1: 7, 2: 5, 3: 6, 4: 1, 7: 7, 8: 9}
	Buyday :=getbuyday(myarray)
	sellday := getsellday(myarray)
	fmt.Printf("The best day to buy is day - %d and the best day to sell is day - %d \n",Buyday, sellday)
}
// get the buy day
func getbuyday(data map[int]int) int {
	// buy day should be at the lowest
	//  to sort array by key
	res := getkeys(data)
	sorteddata := sortincremental(res)
	bestprice := sorteddata[0]
	bestday := getprice(bestprice, data)
	if bestday == len(data)-1{
		res := getkeys(data)
		sorteddata := sortincremental(res)
		sorteddata = sorteddata[:len(sorteddata)-1]
		bestprice := sorteddata[0]
		bestday := getprice(bestprice, data)
		return bestday
	}else{
		return bestday
	}
}
// get the best selling day 
func getsellday(data map[int]int) int {
	// sell day should atleast be 1day after the buy day
	// selling day should at the hieghest
	res := getkeys(data)
	sorteddata := sortincremental(res)
	bestprice := sorteddata[len(sorteddata) -1]
	bestsellday := getprice(bestprice, data)
	return bestsellday
}

// get keys from the map
func getkeys(data map[int]int) []int {
	results := []int{}
	for _, v := range data {
		results = append(results, v)
	}
	return results
}

// sort keys received
func sortincremental(data []int) []int {
	for i:=0; i< len(data); i++{
		for j := i+1; j< len(data); j++{
			if data[i] > data[j]{
				data[i], data[j] = data[j], data[i]
			} 
		}
	}
	return data
}
//get price by looping through the map
func getprice(bestprice int,data map[int]int) int{
	day := 0
	for x,y := range data{
		if y == bestprice {
			day =x
		}
	}
	return day
}
