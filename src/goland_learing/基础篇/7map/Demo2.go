package main

import "fmt"

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 99
	fmt.Println(scoreMap)
	// value, ok := map[key]
	value, ok := scoreMap["张1"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("没有")
	}
}
