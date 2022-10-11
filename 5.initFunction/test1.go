package main

import (
	"./lib"
	t3 "./test3"
	"fmt"
)

func init() {
	fmt.Printf("test1 init.....")
}

func main() {
	fmt.Println("main.......")
	lib.Test44()
	t3.Test3()
	t3.Test22()
}
