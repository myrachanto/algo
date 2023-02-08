package main

import "fmt"

type stack struct{
	data []int
}
func (s *stack) Push(i int){
	s.data = append(s.data, i)
}
func (s *stack) Pop() int{
	lastindex := len(s.data) -1
	s.data = s.data[:lastindex]
	return lastindex
}
func main(){
	mydata := &stack{}
	mydata.Push(100)
	mydata.Push(200)
	mydata.Push(300)
	fmt.Println(mydata)
	mydata.Pop()
	fmt.Println(mydata)


}