package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 3; i++ {
		sum++
	}
	fmt.Println(sum)

	str := "hello"
	fmt.Printf("%s \n", str)
	for i := range str {
		fmt.Println(i, str[i])
		fmt.Printf("%q \n", str[i])
	}
	fmt.Println("********************")
	for a, b := range str {
		fmt.Println(a)
		fmt.Println(b)
		fmt.Printf("%q \n", b)
	}
}
