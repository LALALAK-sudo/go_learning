package main

import "fmt"

// 函数作为 参数传递

func main() {
	fmt.Println(cal(1, 2, add1))
}

func add1(a int, b int) int {
	return a + b
}

// 函数作为参数传递
// 名称 结构 返回值
func cal(a int, b int, op func(int, int) int) int {
	return op(a, b)
}
