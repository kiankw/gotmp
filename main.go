package main

import "fmt"

type MySeeker interface {
	MySeek(name string)
}

type MyFile interface {
	MySeeker
}

type A struct {
	F MyFile
}

func (a A) MySeek(name string) {
	a.F.MySeek(name)
}

type B struct {
	S MySeeker
}

func (b B) MySeek(name string) {
	b.S.MySeek(name)
}

type C struct {
}

func (c C) MySeek(name string) {
	fmt.Println(name)
}

func main() {
	var a A
	var b B
	var c C

	a.F = c
	b.S = c

	a.MySeek("Hello")
	b.MySeek("Hello")

}
