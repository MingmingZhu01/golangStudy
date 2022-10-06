package main

import "fmt"

func main() {
	// 1.声明一个变量默认为0
	var a int
	fmt.Println("a =", a)

	// 2.声明要给变量,直接初始化
	var b int = 100
	fmt.Println("b =", b)

	// 3.声明变量是不用给出具体类型
	var c = 100
	fmt.Println("c =", c)
	fmt.Printf("type of c = %T\n", c)

	// 4.声明一个变量直接省去var关键字,直接自动类型匹配
	d := 100
	f := "abc"
	fmt.Println("d =", d)
	fmt.Printf("type of c = %T\n", d)
	fmt.Println("f =", f)
	fmt.Printf("type of f = %T\n", f)

	// 多变量同时声明
	var e, g int
	e = 100
	g = 100
	fmt.Println("e =,", e, "g =", g)

	var h, i = 100, "abcd"
	fmt.Println("h =,", h, "i =", i)

	var (
		vv int  = 12
		bb bool = true
	)

	fmt.Println("vv =,", vv, "bb =", bb)

	cc, dd := 100, 100
	fmt.Println("cc =,", cc, "bb =", dd)
}
