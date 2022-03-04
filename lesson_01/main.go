package main

import "fmt"

func main() {
	// // boolean
	// var flag bool

	// fmt.Println(flag)

	// // Integer

	// var num int64 = 10
	// fmt.Println(num)

	// // float
	// var num1 float64 = 10.7
	// var num2 float32 = 10.5
	// fmt.Println(num1, num2)

	// var vs :=

	// var x int = 10
	// var x = 10
	// var x int
	// fmt.Println(x)

	// var x, y = 10, "patika"
	// var (
	// 	x    int
	// 	y        = 20
	// 	z    int = 30
	// 	f, g string
	// )

	// const
	// const pi = 3

	// const MAXVALUE = 3
	// const MaxValue = 3

	// Composite Types
	// Arrays

	// var x [3]int
	// var x = [3]int{10, 20, 30}
	// var y = [...]int{10, 20, 30}

	// fmt.Println(x == y) // True

	// matris
	// var x [2][2]int
	// fmt.Println(x)

	//slices
	// [] -> slice
	// [...] -> array

	// x := [3]int{1, 2, 3}
	// fmt.Println(x)

	// append
	// var x []int
	// x = append(x, 10, 20)
	// fmt.Println(x)

	// var x []int
	// y := []int{10, 20, 30}
	// x = append(x, y...)
	// fmt.Println(x)

	// slice len and cap
	// var x []int
	// fmt.Println(x, len(x), cap(x)) //[] 0 0

	// x = append(x, 10)
	// fmt.Println(x, len(x), cap(x)) //[10] 1 1

	// x = append(x, 20)
	// fmt.Println(x, len(x), cap(x)) //[10 20] 2 2

	// x = append(x, 30)
	// fmt.Println(x, len(x), cap(x)) //[10 20] 3 4

	// x = append(x, 20)
	// fmt.Println(x, len(x), cap(x)) //[10 20] 4 4

	// x = append(x, 20)
	// fmt.Println(x, len(x), cap(x)) //[10 20] 5 8

	// make
	// x := make([]int, 3)
	// fmt.Println(x, len(x), cap(x)) //[0 0 0] 3 3

	// x = append(x, 90)
	// fmt.Println(x, len(x), cap(x)) //[0 0 0 90] 4 8

	// x := make([]int, 5, 10)
	// fmt.Println(x, len(x), cap(x)) //[0 0 0 0 0] 5 10

	// Slicing Slice
	// x := []int{1, 2, 3, 4}
	// y := x[:2]
	// fmt.Println(y) // [1 2]

	// e := x[:]
	// fmt.Println(e) // [1 2 3 4]

	// x := [4]int{1, 2, 3, 4}

	// y := x[:3]
	// fmt.Println(y) // [1 2 3]

	// Copy
	// x := []int{1, 2, 3, 4}
	// y := make([]int, len(x))
	// fmt.Println(y)
	// num := copy(y, x)

	// fmt.Println(y, num) // [1 2 3 4] 4

	// Maps

	// var x map[string]int
	// fmt.Println(x) // map[] nil

	// totalWins := map[string]int{}
	// map function yada channel alabilir
	// teams := map[string][]string{
	// 	"Patika":  []string{"a", "d", "d"},
	// 	"Patika1": []string{"a", "d", "d"},
	// 	"Patika2": []string{"a", "d", "d"},
	// }
	// teams["emre"] = []string{"a", "d", "d"}
	// // map[Patika:[a d d] Patika1:[a d d] Patika2:[a d d] emre:[a d d]]
	// fmt.Println(teams)

	// ages := make(map[int][]string, 10)
	// fmt.Println(ages)

	// Reading and writing map

	// totalWins := map[string]int{}
	// totalWins["Patika"] = 10
	// totalWins["Patika1"] = 20
	// totalWins["Patika1"]++
	// fmt.Println(totalWins)
	// fmt.Println(totalWins["emre"]) // 0

	// v, ok := totalWins["emre"]
	// if !ok {
	// 	fmt.Println("emre not found")
	// 	return
	// }
	// fmt.Println(v)

	// struct
	// type person struct {
	// 	name string
	// 	age  int
	// 	pet  string
	// }
	// // var ahmet person
	// // ayse := person

	// osman := person{
	// 	name: "osman",
	// 	age:  30,
	// 	pet:  "dog",
	// }
	// fmt.Println(osman) //{osman 30 dog}
	// osman.age = 21
	// fmt.Println(osman) //{osman 21 dog}

	// anonymous struct
	// bır yerden baska yere pass edılmeycekse kullanılabılır

	// var person struct {
	// 	name string
	// 	age  int
	// 	pet  string
	// }
	// person.name = "Osman"
	// person.age = 30
	// person.pet = "dog"
	// fmt.Println(person)

	// pet := struct {
	// 	name string
	// 	kind string
	// }{
	// 	name: "efe",
	// 	kind: "dog",
	// }

	// type firstPerson struct {
	// 	name string
	// 	age  int
	// }
	// f := firstPerson{
	// 	name: "osman",
	// 	age:  30,
	// }
	// var g struct {
	// 	name string
	// 	age  int
	// }
	// fmt.Printf("%T\n", g) //struct { name string; age int }
	// g.name = "Osman"
	// g.age = 31
	// g = f
	// fmt.Println(f == g) // true

	// Shadowing variable

	// x := 10

	// if x > 5 {
	// 	fmt.Println(x) //10
	// 	x := 5
	// 	fmt.Println(x) //5
	// }
	// fmt.Println(x) //10

	// x := 10

	// if x > 5 {
	// 	fmt.Println(x) //10
	// 	x = 5
	// 	fmt.Println(x) //5
	// }
	// fmt.Println(x) //5

	// if statement

	// n := 10

	// if n == 0 {
	// 	//
	// } else if n == 1 {
	// 	//
	// } else {
	// 	//
	// }

	// for loop
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// i := 1
	// for i < 100 {
	// 	fmt.Println(i)
	// 	i = i * 2
	// }

	// evenVals := []int{2, 4, 6, 8}
	// for i, v := range evenVals {
	// 	fmt.Println(i, v)
	// }
	//0 2
	// 1 4
	// 2 6
	// 3 8

	// mapsler hash tabledan dolayı for döngüsünde farklı sırada gelir

	// switch case

	words := []string{"dog", "cat", "bird", "fish"}

	for _, word := range words {
		switch word {
		case "dog":
			fmt.Println("dog")
		case "cat":
			fmt.Println("cat")
		}

	}
}
