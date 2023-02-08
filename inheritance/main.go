package main

import "fmt"

type Node struct {
	Key   int
	Value string
	*Node
}
type Modes struct {
	Key   int
	Value string
	Mode
}
type Mode struct {
	Key   int
	Value string
}

func main() {
	node := Node{
		Key:   1,
		Value: "first",
		Node: &Node{
			Key:   2,
			Value: "Second",
			Node: &Node{
				Key:   3,
				Value: "Third",
				Node: &Node{
					Key:   4,
					Value: "fourth",
					Node:  &Node{}},
			},
		},
	}
	fmt.Println(node)
	fmt.Println(*node.Node.Node.Node)
	//		mode := Modes{
	//			Key:   1,
	//			Value: "first",
	//			Mode: Mode{
	//				Key:   2,
	//				Value: "Second",
	//			},
	//		}
	//		fmt.Println(mode)
	//	}
}
