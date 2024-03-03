package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	os.Setenv("TZ", "Asia/Tokyo")

	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Hour())

	// 条件文がないswitch文も書ける
	// trueであることを示す
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}
}