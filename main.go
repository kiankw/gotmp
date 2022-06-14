package main

type A interface {
	fa(name string)
}

type B interface {
	A
	fb(name string)
}

type C interface {
	A
	fc(name string)
}

type D interface {
	B
	C
	fd(name string)
}

type T struct {
	id int
}

func (t T) fa(name string) {

}

func (t T) fb(name string) {

}

func (t T) fc(name string) {

}

func (t T) fd(name string) {

}

func main() {
	var a A
	a = T{id: 1}
	a.fa("Hello")
}
