// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {

// 	var number int

// 	//r := rand.Intn(10) + 1

// 	var res bool = true

// 	for i := 1; res == true; i++ {

// 		s := rand.NewSource(time.Now().UnixNano())
// 		r := rand.New(s)
// 		//fmt.Println("R", r)
// 		//fmt.Println("S", s)
// 		result := r.Intn(10)

// 		fmt.Println("Guess the number?")
// 		fmt.Scanln(&number)

// 		if result == number {
// 			//fmt.Println("Random no is", result)
// 			fmt.Println("Yay! You got it!!!")
// 			res = false
// 			break

// 		} else if i == 3 {
// 			fmt.Println("Game Over!")
// 			res = false
// 			break
// 		} else {

// 			fmt.Println("Bad Luck, Try again")
// 			//fmt.Println("Random no is", result)
// 			continue
// 		}

// 	}

// }
