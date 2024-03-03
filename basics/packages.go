// Goプログラムは、パッケージで構成される

// プログラムはmainパッケージから開始
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Intn()は10以下の疑似乱数を返す
	fmt.Println("My favorite number is", rand.Intn(10))
}