// package main

// import (
// 	"fmt"
// )

// func main() {

// 	s := make([]string, 3)

// 	fmt.Println("Original Slice:", s)

// 	s[0] = "a"
// 	s[1] = "b"
// 	s[2] = "c"

// 	fmt.Println("Slice s contains:", s)

// 	fmt.Println("Element at 3rd position:", s[2])

// 	fmt.Println("length of the slice:", len(s))

// 	s = append(s, "d")
// 	s = append(s, "e", "f")

// 	fmt.Println("New slice becomes: ", s)

// 	c := make([]string, len(s))
// 	copy(c, s)

// 	fmt.Println("Slice C is: ", c)

// 	l := s[2:5]

// 	fmt.Println("Slice between 2 to 5 are:", l)

// 	l1 := s[:3]

// 	fmt.Println("Slice upto 3 :", l1)

// 	l2 := s[2:]

// 	fmt.Println("Slice after 2: ", l2)

// 	t := []string{"x", "y", "z"}

// 	fmt.Println("Slice t is:", t)

// 	fmt.Println("Range of slice t is:", len(t))

// }
