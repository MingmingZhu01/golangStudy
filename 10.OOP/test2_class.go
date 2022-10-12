package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}

//func (this Hero) Show() {
//	fmt.Println("Name = ", this.Name)
//	fmt.Println("Ad = ", this.Ad)
//	fmt.Println("Level = ", this.Level)
//}
//
//func (this Hero) GetName() string {
//	return this.Name
//}
//
//func (this Hero) SetName(newName string) {
//	//this 是调用该方法的对象的一个副本（拷贝）
//	this.Name = newName
//}

func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	//this 是调用该方法的对象的一个副本（拷贝）
	this.Name = newName
}

func main() {
	//创建一个对象
	hero := Hero{Name: "zhang3", Ad: 100}

	hero.Show()

	hero.SetName("li4")

	hero.Show()
}
