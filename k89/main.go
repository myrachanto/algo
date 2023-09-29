package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const (
	name     = "K89: "
	address  = "127.0.0.10"
	address1 = "127.0.0.1"
)

type (
	SystemControllerInterface interface {
		StopSystem()
		SystemStatus()
		AddNode(string) error
		DeleteNode(*Node) (bool, error)
		SchedulePod(*Pod) (bool, error)
		DeletePod(*Pod) (bool, error)
		PodStatus(*Pod) (bool, error)
	}
	NodeInterface interface {
		StartBackground() (bool, error)
	}
)
type (
	SystemController struct {
		Name    string
		Nodes   map[string]*Node
		Address string
		Status  bool
	}
	Node struct {
		Name        string
		Pods        map[string]*Pod
		Address     string
		Status      bool
		Schedulable bool
	}
	Pod struct {
		Name        string
		Image       string
		Schedulable bool
		Status      bool
		Ports       bool
		StartTime   time.Time
	}
)

func New() *SystemController {
	log.SetPrefix(name)
	log.Println("Started.......")
	return &SystemController{
		Name:    name,
		Address: address,
		Nodes:   make(map[string]*Node),
		Status:  true,
	}
}
func (sc *SystemController) StopSystem() {
	sc.Status = false
	log.Println("Stoped!")
}
func (sc *SystemController) SystemStatus() {
	if sc.Status {
		log.Println("is running!...")
	} else {
		log.Println("has being Stoped!")
	}
}
func (sc *SystemController) AddNode(name string) error {
	schedulable := IsSchedulable(name)
	node := sc.NewNode(name, schedulable)
	err := node.StartBackground()
	if err != nil {
		return err
	} else {
		sc.Nodes[node.Name] = node
		return nil
	}
}
func (sc *SystemController) NewNode(name string, schedulable bool) *Node {
	if name == "" {
		log.Panic("Please use a valid Name for a Node!")
	}
	address, err := sc.CreateAddress(name)
	if err != nil {
		log.Panic("something Went wrong with creating an Address")
	}
	return &Node{
		Name:        name,
		Pods:        make(map[string]*Pod),
		Address:     address,
		Schedulable: schedulable,
	}
}
func (n *Node) StartBackground() error {
	n.Status = true
	log.Printf("%s background services started \n", n.Name)
	return nil
}
func (sc *SystemController) CreateAddress(name string) (string, error) {
	if len(sc.Nodes) == 0 {
		return address1, nil
	} else {
		l := len(sc.Nodes)
		return fmt.Sprintf("127.0.0.%d", l), nil
	}
}
func IsSchedulable(name string) bool {
	res := false
	lastString := strings.Split(name, "_")
	lastStringSchedulable := lastString[len(lastString)-1]
	if lastStringSchedulable == "master" {
		res = true
		return res
	}
	return res
}

func main() {

}
