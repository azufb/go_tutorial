package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	// 構造体vのメモリアドレスをpに代入
	p := &v
	fmt.Println(*p)// {1 2}

	p.X = 1e9
	fmt.Println(*p)

	// ポインタに入っているメモリアドレスにアクセスし、そこに100を代入
	// アスタリスクなしで、p.Xと記述する方が手軽
	(*p).X = 100
	fmt.Println(*p) // {100 2}

	fmt.Println(v) // {100, 2}
}