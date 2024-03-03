package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// sliceは配列より一般的で、可変長
	// 配列名[始めのindex:終わりのindex(含まれない)]で、始めのindex～終わりのindex-1の値を要素とする
	var s []int = primes[1:4]
	fmt.Println(s)
}