package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 3, 2, 1, 3, 6, 3, 2}
	sort.Ints(a)
	fmt.Println(a)
}
