package main

import "fmt"

func printArray(myArray []int) {
	//引用传递
	// _ 表示匿名的变量
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
}

func main() {
	myArray := []int{1, 23}
	var myArray1 []int
	fmt.Printf("myArray type is %T\n", myArray)
	fmt.Printf("myArray1 type is %T\n", myArray1)

	printArray(myArray)

	fmt.Println(" ==== ")

	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
	fmt.Println("===========")
	myArray1 = append(myArray1, 1, 2, 3)
	printArray(myArray1)

	fmt.Println("===========")
	slice1 := make([]int, 0)
	slice1 = append(slice1, 1, 1, 1)
	printArray(slice1)

	fmt.Println("===========")
	s2 := slice1[0:2]
	printArray(s2)

	fmt.Println("===========")
	s3 := []int{1, 2, 3}
	copy(s3, s2)
	fmt.Println(s2)
	fmt.Println(s3)
}
