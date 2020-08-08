// package main

// import "fmt"

// func main() {

// 	//declaring a slice of strings
// 	s := []string{"a", "b", "c", "d"}

// 	// Here we use range to check each element in slice
// 	for i, c := range s {
// 		fmt.Println("Index:", i)
// 		fmt.Println("C is: ", c)
// 	}

// 	fmt.Println()
// 	fmt.Println()
// 	//declaring a map of key value pairs
// 	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}

// 	//Here we use range to check each value pair element
// 	for k, v := range m {
// 		fmt.Println("Keys:", k)
// 		fmt.Println("Values:", v)
// 	}

// 	fmt.Println()
// 	fmt.Println()

// 	//check element of above mentioned slice with _ operator
// 	for _, c := range s {
// 		fmt.Println("C is:", c)
// 	}

// 	//adding the int elements of slice using range
// 	nums := []int{1, 2, 3, 4, 5}
// 	sum := 0
// 	for _, num := range nums {
// 		sum = sum + num
// 	}
// 	fmt.Println("Sum is:", sum)
// }
