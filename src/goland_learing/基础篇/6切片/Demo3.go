package main

import "fmt"

func main() {
	// append()添加元素和切片扩容
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v len:%d cap:%d prt:%p \n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}
