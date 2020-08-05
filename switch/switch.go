// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// 	i := 3
// 	switch i {
// 	case 1:
// 		fmt.Println(i, "is one")

// 	case 2:
// 		fmt.Println(i, "is two")

// 	case 3:
// 		fmt.Println(i, "is three")
// 	}

// 	switch time.Now().Weekday() {
// 	case time.Saturday, time.Sunday:
// 		fmt.Println("It is weekend")

// 	default:
// 		fmt.Println(time.Now().Weekday())
// 		fmt.Println("It is weekday")

// 	}

// 	t := time.Now()
// 	switch {
// 	case t.Hour() < 12:
// 		fmt.Println("It is morning")
// 	case t.Hour() > 12:
// 		fmt.Println("It is noon")

// 	}

// }
