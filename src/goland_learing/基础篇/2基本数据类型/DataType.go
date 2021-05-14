package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 10
	var b int = 077
	var c int = 0xff
	fmt.Println(a, b, c)
	fmt.Printf("%d \n", a)
	fmt.Printf("%b \n", a)

	fmt.Printf("%o \n", b)
	fmt.Printf("%x \n", c)
	fmt.Printf("%f \n", math.Pi)
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")
	s1 := `第一行
	第二行
	第三行
	`
	fmt.Println(s1)

	var s = "hello"
	str := " string"
	sprint := fmt.Sprint(s, str)
	fmt.Println(sprint)

	changeString()
}

func changeString() {
	s1 := "very handsome"
	bytes := []byte(s1)
	bytes[0] = 'i'
	fmt.Println(string(bytes))

	s2 := "我很帅"
	runes := []rune(s2)
	runes[2] = '美'
	fmt.Println(string(runes))
}
