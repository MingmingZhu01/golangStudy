package main

import "fmt"

// 1.无参数,无返回值
func foo1() {
	fmt.Println("foo1..........")
}

// 2.无参数,有返回值
func foo2() int {
	fmt.Println("foo2..........")
	return 1
}

// 3.有参数,无返回值

func foo3(x int) {
	fmt.Println("foo3..........", x)
}

// 4.有参数,有返回值
func foo4(x int) int {
	fmt.Println("foo4..........", x)
	return x + 1
}

// 5.多返回值,匿名
func foo5(x int) (int, int) {
	fmt.Println("foo5........")
	return 100, 200
}

// 6.多返回值,非匿名

func foo6(x int, y int) (r1 int, r2 int) {
	fmt.Println("foo6........")
	r1 = 1000
	r2 = 2000
	return // 可不写return r1, r2 已经默认返回r1, r2。 不初始化r1,r2的话,默认值为0
}

func main() {
	foo1()
	foo2()
	foo3(111)
	a := foo4(1)
	fmt.Println("a =", a)
	b, c := foo5(1)
	fmt.Println("b =", b, ", c = ", c)
	d, f := foo6(1, 2)
	fmt.Println("d =", d, ", f = ", f)
}
