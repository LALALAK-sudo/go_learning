package main

import "fmt"

func main() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}

	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
}
