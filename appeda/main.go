package main

import "fmt"

func main() {
	list := make([]int, 0, 10)
	list = append(list, 1, 2, 3, 4, 5)
	fmt.Printf("the length is %d and capcity is %d \n", len(list), cap(list))
	appender(list)
	fmt.Println(">>>>>>>>slice", list)

	arr := map[string]int{"tony": 30}
	fmt.Println(">>>>>>>>arr", arr)
	mapper(arr)
	fmt.Println(">>>>>>>>arr", arr)
}

func appender(data []int) {
	for k := range data {
		data[k] += 10
	}
	data = append(data, 16, 17, 18, 19, 20)
	fmt.Println("data-----", data)

}
func mapper(data map[string]int) {
	data["tony"] = 32
	data["jenny"] = 25
}
