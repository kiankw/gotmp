package main

import "fmt"

type A struct {
	info func() string
}

var ATable = []*A{{
	info: f0,
}, {
	info: f1,
}}

func f0() string {
	return "func 0"
}

func f1() string {
	return "func 1"
}

func main() {
	var form int
	fmt.Scanln(&form)

	a := *ATable[form]
	fmt.Println(a.info())
}
