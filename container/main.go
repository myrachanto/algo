package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func init() {
	log.SetPrefix("Container: ")
}

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		log.Panic("something is very wrong")
	}

}
func run() {
	fmt.Printf("Running %v on %d\n", os.Args[2:], os.Getpid())
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	must(cmd.Run())
}
func must(err error) {
	if err != nil {
		log.Panic(err)
	}
}
