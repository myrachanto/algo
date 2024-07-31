package main

// import "fmt"

type number struct {
	price float64
}

func (n *number) destroy() {
	// fmt.Println("============ boom")
	n = nil
}
func main() {
	n := number{45}
	// fmt.Println("===========", n)
	n.destroy()
	// fmt.Println("===========", n)
	escapeheap()
	// fmt.Println("===========", *a)
}
func escapeheap() *int {

	a := 42
	return &a
}
