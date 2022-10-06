package main

import "fmt"

const v1 int = 1000
const (
	v2 = 1000
	v3 = 1000
)

const (
	// BEIJING 可以在const() 添加一个关键字 iota， 每行的iota都会累加1, 第一行的iota的默认值是0
	BEIJING  = 10 * iota //iota = 0
	SHANGHAI             //iota = 1
	SHENZHEN             //iota = 2
)

const (
	a, b = iota + 1, iota + 2 // iota = 0, a = iota + 1, b = iota + 2, a = 1, b = 2
	c, d                      // iota = 1, c = iota + 1, d = iota + 2, c = 2, d = 3
	e, f                      // iota = 2, e = iota + 1, f = iota + 2, e = 3, f = 4

	g, h = iota * 2, iota * 3 // iota = 3, g = iota * 2, h = iota * 3, g = 6, h = 9
	i, k                      // iota = 4, i = iota * 2, k = iota * 3 , i = 8, k = 12
)

func main() {
	fmt.Println("v1 =", v1)
	fmt.Println("v2 =", v2)
	fmt.Println("v3 =", v3)

	fmt.Println("BEIJIGN = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)

	fmt.Println("g = ", g, "h = ", h)
}
