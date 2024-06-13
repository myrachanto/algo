package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/google/uuid"
)

func main() {
	// Generate a new UUID
	newUUID := uuid.New()

	// Convert the UUID to a string
	uuidStr := newUUID.String()

	// Remove the hyphens from the UUID string
	uuidStrNoHyphens := ""
	for _, char := range uuidStr {
		if char != '-' {
			uuidStrNoHyphens += string(char)
		}
	}

	// Convert the hexadecimal string to a big integer
	uuidInt := new(big.Int)
	uuidInt.SetString(uuidStrNoHyphens, 16)

	// Print the results
	fmt.Println("UUID as a string:", uuidStr)
	fmt.Println("UUID as an integer:", uuidInt)

	uuid := "629a7b60-89d8-433a-94d0-015b9cb69a99"

	// Remove dashes from UUID
	uuid = strings.ReplaceAll(uuid, "-", "")

	// Compute MD5 hash
	hash := md5.Sum([]byte(uuid))

	// Convert hash to hexadecimal string
	hashString := hex.EncodeToString(hash[:])

	// Convert hexadecimal to decimal
	intVal, _ := new(big.Int).SetString(hashString, 16)

	// Convert to 13-digit code
	code := intVal.Mod(intVal, big.NewInt(10000000000000)).String()

	fmt.Println("13-digit code:", code)
}
