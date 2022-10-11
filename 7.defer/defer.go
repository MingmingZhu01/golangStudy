package main

import "fmt"

func main() {
	// defer关键字
	defer fmt.Println("defer1.......")
	defer fmt.Println("defer2.......")

	fmt.Println("aaaaa")
	fmt.Println("bbbbbb")
}
