package main

import (
	"fmt"
	"math"
)

type MyFloat float64
type Name string

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (n Name) GiveMessage() string {
	// Name型になっているので型変換
	nameStr := string(n)
	message := nameStr + "さん、こんにちは！"
	return message
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	n := Name("Azusa")
	fmt.Println(n.GiveMessage())
}
