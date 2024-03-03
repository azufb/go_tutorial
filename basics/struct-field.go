package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// フィールドには、ドットでアクセス
	fmt.Println("X:before", v.X)
	fmt.Println("Y:before", v.Y)

	v.X = 4
	v.Y = 100
	fmt.Println("X:after", v.X)
	fmt.Println("Y:after", v.Y)
}