// package main

// import "fmt"

// func zeroVal(inVal int) {
// 	inVal = 0
// }

// func zeroPtr(inPtr *int) {
// 	*inPtr = 0
// }

// func main() {

// 	i := 1
// 	fmt.Println("Initial:", i)

// 	zeroVal(i)
// 	fmt.Println("zeroVal:", i) //The value of i remains the same

// 	zeroPtr(&i)
// 	fmt.Println("zeroPtr:", i) // The value of i changes because the refrence is set to 0

// 	fmt.Println("Pointer:", &i) // Printing the pointer value as well
// }
