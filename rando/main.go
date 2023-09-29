package main

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alfabet = "abcdefghijklmnopqrstuvwxyz"
	pasword = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789?$@!&-_"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generate random number between max and min
func RandomInt(max, min int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Generate a random string
func RandomString(n int) string {
	var str strings.Builder
	l := len(alfabet)
	for i := 0; i < n; i++ {
		c := alfabet[rand.Intn(l)]
		str.WriteByte(c)
	}
	return str.String()
}
func RandomPassword(n int) string {
	var pass strings.Builder
	k := len(pasword)
	for i := 0; i < n; i++ {
		c := pasword[rand.Intn(k)]
		pass.WriteByte(c)
	}
	return pass.String()

}

func main() {

}
