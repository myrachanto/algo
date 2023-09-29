package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		words[scan.Text()]++

	}
	fmt.Fprintln(os.Stdout, len(words), "unique words")
	type kv struct {
		key   string
		value int
	}
	var res []kv
	for a, b := range words {
		res = append(res, kv{a, b})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].value > res[j].value
	})
	for _, g := range res {
		fmt.Println(g.key, g.value)
	}
}
