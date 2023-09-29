package main

import (
	"fmt"
)

// Node represents a key-value pair in the hash map.
type Node struct {
	Key   string
	Value int
}

// HashMap represents the hash map data structure.
type HashMap struct {
	buckets []*Node
	size    int
}

// NewHashMap creates a new hash map with the given size.
func NewHashMap(size int) *HashMap {
	return &HashMap{
		buckets: make([]*Node, size),
		size:    size,
	}
}

// hashFunction calculates the index for a given key.
func (hm *HashMap) hashFunction(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + int(key[i])) % hm.size
	}
	return hash
}

// Insert inserts a key-value pair into the hash map.
func (hm *HashMap) Insert(key string, value int) {
	index := hm.hashFunction(key)
	hm.buckets[index] = &Node{Key: key, Value: value}
}

// Get retrieves the value associated with a given key.
func (hm *HashMap) Get(key string) (int, bool) {
	index := hm.hashFunction(key)
	node := hm.buckets[index]
	if node == nil {
		return 0, false
	}
	return node.Value, true
}

// Delete removes a key-value pair from the hash map.
func (hm *HashMap) Delete(key string) {
	index := hm.hashFunction(key)
	hm.buckets[index] = nil
}

func main() {
	hashMap := NewHashMap(10)

	// Insert some key-value pairs
	hashMap.Insert("Alice", 25)
	hashMap.Insert("Bob", 30)
	hashMap.Insert("Charlie", 35)

	// Retrieve values by key
	if value, exists := hashMap.Get("Alice"); exists {
		fmt.Println("Alice's age is", value)
	}

	// Delete a key-value pair
	hashMap.Delete("Bob")

	// Try to retrieve the deleted key
	if _, exists := hashMap.Get("Bob"); !exists {
		fmt.Println("Bob's information is not found")
	}
}

// import (
// 	"crypto/md5"
// 	"fmt"
// 	"reflect"
// )

// type Hashtable[T comparable, K comparable] struct {
// 	Buckets []*items[T, K]
// 	Size    int
// }
// type items[T comparable, K comparable] struct {
// 	Name  string
// 	Store []item[T, K]
// }
// type item[T comparable, K comparable] struct {
// 	name  T
// 	value K
// }

// //	func NewHastable(size ...int) *Hashtable {
// //		sizer := 0
// //		if size[0] == 0 {
// //			sizer = 1000
// //		}
// //		return &Hashtable{
// //			Buckets: make([]*items[T, K], sizer),
// //		}
// //	}
// func (hm *Hashtable[T, K]) hashFunction(key string) int {
// 	hash := 0
// 	for i := 0; i < len(key); i++ {
// 		hash = (hash + int(key[i])) % hm.Size
// 	}
// 	return hash
// }
// func (h *Hashtable[T, K]) put(key T, value K) {
// 	hash := md5.New()
// 	hash.Write(converter[T](key))
// }

// func converter[T comparable](key T) []byte {
// 	if reflect.TypeOf(key).Kind() == reflect.String {
// 		return []byte(fmt.Sprintf("%s", key))
// 	}
// 	if reflect.TypeOf(key).Kind() == reflect.Int {
// 		return []byte(fmt.Sprintf("%d", key))
// 	}
// 	if reflect.TypeOf(key).Kind() == reflect.Int16 {
// 		return []byte(fmt.Sprintf("%d", key))
// 	}
// 	if reflect.TypeOf(key).Kind() == reflect.Int32 {
// 		return []byte(fmt.Sprintf("%d", key))
// 	}
// 	if reflect.TypeOf(key).Kind() == reflect.Int64 {
// 		return []byte(fmt.Sprintf("%d", key))
// 	}
// 	if reflect.TypeOf(key).Kind() == reflect.Float32 {
// 		return []byte(fmt.Sprintf("%2f", key))
// 	}
// 	if reflect.TypeOf(key).Kind() == reflect.Float64 {
// 		return []byte(fmt.Sprintf("%2f", key))
// 	}
// 	return []byte{}
// }
