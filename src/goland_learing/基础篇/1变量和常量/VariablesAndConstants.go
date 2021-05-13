package main

import "fmt"

// 全局变量 a
var a = 100

func main() {
	// 标准声明1
	// var 变量名 变量类型
	var name string
	//fmt.Printf(name)
	name = "hello"
	fmt.Println(name)

	// 标准声明2
	// var 变量名 类型 = 表达式
	var name2 string = "你好"
	fmt.Println(name2)

	// 标准声明3
	var name3 string = "自动推导了"
	fmt.Println(name3)

	fmt.Println("全局变量", a)

	// 标准声明4
	// 短变量声明
	a := 200
	b := 1
	fmt.Println("局部变量", a)
	fmt.Println("局部变量", b)

	// 标准声明5
	// 匿名变量
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)

}
func foo() (int, string) {
	return 10, "Q1mi"
}
