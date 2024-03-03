package main

import "fmt"

func main() {
	// deferすることで、関数が最後まで実行されてreturnされるまで待つ
	// returnされたらdeferされたコードが実行される
	defer fmt.Println("World!")

	fmt.Println("Hello")
}