package main

import "fmt"

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {
	a := 20
	b := 10
	swap(&a, &b)
	fmt.Println("a =", a, "b =", b)

	var p *int

	p = &a

	fmt.Println(&a)
	fmt.Println(p)

	var pp **int //二级指针

	pp = &p

	fmt.Println(&p)
	fmt.Println(pp)
}
