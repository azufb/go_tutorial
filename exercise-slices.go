package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// 長さdyのslice
	s := make(uint8, dy)

	x := make(uint8, dx)
}

func main() {
	pic.Show(Pic)
}