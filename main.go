package main

import "fmt"

// type MyFunc func(n int) string

// type A struct {
// 	info func() string
// }

// var ATable = []*A{{
// 	info: f0,
// }, {
// 	info: f1,
// }}

// func f0() string {
// 	return "func 0"
// }

// func f1() string {
// 	return "func 1"
// }

// func main1() {
// 	var farr [5]MyFunc
// 	farr[0] = func(n int) string {
// 		fmt.Println("func0")
// 		return "func0"
// 	}
// 	var form int
// 	fmt.Scanln(&form)
// 	farr[0](form)

// 	a := *ATable[form]
// 	fmt.Println(a.info())
// }

func A() string {
	print("A")
	return "A"
}
func B() string {
	print("B")
	return "B"
}

type MyFunc func() string
type S struct {
	f MyFunc
}

func P(s S) {
	s.f()
}

func main() {
	var sArr [2]S
	sArr[0] = S{f: A}
	sArr[1] = S{f: B}

	var n int
	fmt.Scanln(&n)
	P(sArr[n])
}
