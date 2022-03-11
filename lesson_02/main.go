// package main

// import "fmt"

// func noPointer() string {
// 	return "string"
// }

// func pointerTest() *string {
// 	return nil
// }

// func pointerTest2() *string {
// 	s := "string"
// 	return &s
// }

// func main() {
// 	fmt.Println(noPointer())
// 	fmt.Println(pointerTest())
// 	fmt.Println(pointerTest2())

// 	s := pointerTest2()
// 	fmt.Println(s)
// 	sp := *s

// 	fmt.Println(sp)
// }

package main

import (
	"fmt"
)

type User struct {
	Name string
	Pets []string
}

func (u User) newPet() {
	u.Pets = append(u.Pets, "Lucy")
	fmt.Println(u)
}
func main() {
	u := User{Name: "Anna", Pets: []string{"Bailey"}}
	u.newPet()     // {Anna [Bailey Lucy]} -- the user with a new pet, Lucy!
	fmt.Println(u) // Hey, wait! Where did Lucy go?
}
