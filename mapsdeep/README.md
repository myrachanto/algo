## Deep understanding of Golang Maps

# Question 1
Explain why a map cannot have a slice as a key in Go. Then, write a function that attempts to use a slice as a key in a map and demonstrate the compilation error.

# Answer:

In Go, map keys must be of a type that is comparable using the == operator. Slices, being reference types, are not comparable (except comparing to nil), and thus cannot be used as map keys. This prevents the use of slices as keys in maps to avoid unpredictable behavior and runtime errors.