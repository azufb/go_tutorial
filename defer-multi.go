package main

import "fmt"

func main() {
	fmt.Println("Counting...")

	// deferが複数ある場合は、スタックに積まれ、LIFOで取り出される
	for i := 1; i <= 10; i++ {
		defer fmt.Println(i)
	}
	
	fmt.Println("DONE!!")
}