// package main

// import "fmt"

// type person struct {
// 	name string
// 	age  int
// }

// func newPerson(name string) *person {
// 	p := person{name: name}
// 	p.age = 48

// 	return &p

// }

// func main() {

// 	fmt.Println(person{"Bharat", 30})
// 	fmt.Println(person{name: "Yogi", age: 32})
// 	fmt.Println(&person{name: "Sanju"})
// 	fmt.Println(newPerson("Raj"))

// 	s := person{name: "Babbar", age: 44}
// 	fmt.Println(s.name)
// 	fmt.Println(s.age)

// 	sp := &s
// 	fmt.Println(sp.age)
// 	fmt.Println(&sp.age)

// 	sp.age = 45
// 	fmt.Println(sp.age)
// 	fmt.Println(&sp.age)
// }
