package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time // time.Time型は、Goの標準パッケージtimeに用意されている構造体
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// run()関数を実行してnilじゃなかったら、標準出力
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
