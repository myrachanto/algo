package main

import (
	"fmt"
	"os"
)

func Printing(args ...any) {
	for _, v := range args {
		switch x := v.(type) {
		case string:
			os.Stdout.WriteString(x)
		case int:
			os.Stdout.WriteString(String(x))
		case float64:
			os.Stdout.WriteString(String(x))
		}
	}
}
func String(a any) string {
	switch x := a.(type) {
	case float64:
		return fmt.Sprintf(" %.2f ", x)
	case float32:
		return fmt.Sprintf(" %.2f ", x)
	case int:
		return Itoa(x)
	case int32:
		return fmt.Sprintf(" %d ", x)
	case int64:
		return fmt.Sprintf(" %d ", x)
	}
	return ""
}
func itoa(i int) (s string) {
	lookup := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	if i < 0 {
		s = "-"
		i = i * -1
	}
	t := " "
	for i > 0 {
		r := i % 10
		i = i / 10
		t = lookup[r] + t
	}
	return s + t
}

func Itoa(n int) (s string) {
	ltz := n < 0
	if ltz {
		n = -n
	}
	var b [30]byte
	c := len(b) - 1
	for {
		b[c], n, c = byte('0'+n%10), n/10, c-1
		if n == 0 {
			break
		}
	}
	if ltz {
		b[c] = '-'
		c--
	}
	return string(b[c+1:])
}
func main() {
	Printing("vamos a bailar", 123.34, 45, "here we go")
	fmt.Println("vamos a bailar", 123.34, 45, "here we go")
}
