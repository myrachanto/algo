package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, filename := range os.Args[1:] {
		var lc, wc, cc int
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintln(os.Stdout, err)
			continue
		}
		scan := bufio.NewScanner(file)
		for scan.Scan() {
			s := scan.Text()
			wc += len(strings.Fields(s))
			cc += len(s)
			lc++
		}
		fmt.Printf("%7d %7d %7d %5s \n", wc, cc, lc, filename)
		file.Close()
	}
}
