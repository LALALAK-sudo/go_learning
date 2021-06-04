package main

import "fmt"

// 闭包=函数+引用环境
func main() {
	var f = adder()
	fmt.Println(f(10))
	fmt.Println(f(10))
	fmt.Println(f(10))
	fmt.Println(f(10))

}

func adder() func(int) int {
	var x = 10
	return func(y int) int {
		x += y
		return x
	}
}
