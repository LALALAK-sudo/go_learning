package main

import "fmt"

// 可变参数（类似重载，但是更加灵活了）

func main() {
	fmt.Println(sumIntN(1, 2, 3))
}

func sumIntN(x ...int) int {
	fmt.Println(x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
