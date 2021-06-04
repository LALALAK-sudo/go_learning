package main

import "fmt"

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 99
	for key, value := range scoreMap {
		fmt.Println(key, value)
	}
	for a := range scoreMap {
		fmt.Println(a)
	}
	fmt.Println(len(scoreMap))

}
