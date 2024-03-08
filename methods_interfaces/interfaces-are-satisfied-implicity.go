package main

import "fmt"

type I interface {
	// Mというメソッドを実装する
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}
