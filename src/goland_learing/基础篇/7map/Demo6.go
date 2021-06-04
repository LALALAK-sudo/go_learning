package main

import "fmt"

func main() {
	str := "how do you do"
	fmt.Println(str)

	sliceMap := make(map[string]int, 3)
	for i, s := range str {
		fmt.Println(i, s)
		v, ok := sliceMap[string(s)]
		if ok {
			sliceMap[string(s)] = v + 1
		} else {
			sliceMap[string(s)] = 1
		}
	}
	fmt.Println(sliceMap)
}
