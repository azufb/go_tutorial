package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	// 片方のフィールドだけ代入すると、もう片方は0
	v2 = Vertex{X: 1}
	// どちらのフィールドにも値を入れないと、両方0
	v3 = Vertex{}
	p = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}