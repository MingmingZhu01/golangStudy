package main

import "fmt"

func ChangeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	// 第一种
	var map1 map[string]string
	if map1 == nil {
		fmt.Println("map1 是空map")
	}

	// 使用map之前, 需要先用make给map分配数据空间
	map1 = make(map[string]string, 1)
	map1["a"] = "aaa"
	fmt.Println(map1)
	map1["b"] = "bbb"
	fmt.Println(map1)

	// 第二种
	map2 := make(map[string]string)
	map2["c"] = "ccc"
	fmt.Println(map2)

	// 第三种
	map3 := map[string]string{
		"one": "php",
		"two": "java",
	}
	fmt.Println(map3)

	map3["three"] = "go"
	fmt.Println(map3)

	ChangeValue(map3)
	fmt.Println(map3)
	delete(map3, "England")
	fmt.Println(map3)
}
