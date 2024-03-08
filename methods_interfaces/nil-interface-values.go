package main

import "fmt"

type I interface {
	M()
}

type MyString string

func (s MyString) M() {
	fmt.Println(s)
}

func main() {
	var i I
	describe(i)
	// i.M()

	str := MyString("Hello World!")
	i = str
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}