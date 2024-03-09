package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// スレッド
	// go 関数(引数x, 引数y, 引数z)と記述
	go say("world")
	say("hello")
}