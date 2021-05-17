package main

import "fmt"

func main() {
	a := make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		//a = append(a, i)
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}
