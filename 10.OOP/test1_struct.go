package main

import "fmt"

type Book struct {
	title string
	name  string
}

func changeBook(book Book) {
	//传递一个book的副本
	book.name = "666"
}

func changeBook2(book *Book) {
	book.name = "zoe"
}

func main() {

	var book1 Book
	book1.title = "go"
	book1.name = "jack"

	fmt.Printf("%v\n", book1)
	fmt.Println(book1)

	changeBook(book1)
	fmt.Println(book1)

	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
}
