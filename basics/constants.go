package main

import "fmt"

// constで定数宣言
// 文字(character)、文字列(string)、boolean、numericのみ使える
// :=は使えない
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}