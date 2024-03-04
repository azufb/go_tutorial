package main

import "fmt"

func main() {
	// make関数で動的サイズの配列を作る
	// サイズが5、要素がint型の配列
	a := make([]int, 5)
	printSlice("a", a)

	// サイズが0、スライスの容量が5、要素がint型の配列
	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}