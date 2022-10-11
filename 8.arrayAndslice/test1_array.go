package main

import "fmt"

func printArray(myArray [4]int) {
	//值拷贝

	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}

	myArray[0] = 111
}

func main() {
	// 数组定义
	// 1.var 关键字 固定长度数组
	var arr1 [10]int

	// 2.简写
	arr2 := [10]int{1, 1}
	arr3 := [4]int{1, 2, 3, 14}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for index, value := range arr2 {
		fmt.Println(index, " : ", value)
	}
	//查看数组的数据类型
	fmt.Printf("arr1 types = %T\n", arr1)
	fmt.Printf("arr2 types = %T\n", arr2)
	fmt.Printf("arr3 types = %T\n", arr3)

	printArray(arr3)

	fmt.Println(" ------ ")
	for index, value := range arr3 {
		fmt.Println("index = ", index, ", value = ", value)
	}

}
