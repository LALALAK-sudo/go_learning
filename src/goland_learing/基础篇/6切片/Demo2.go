package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 6}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	s2 := s[3:4] // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))

	c := [5]int{1, 2, 3, 4, 5}
	b := c[1:3:5]
	fmt.Printf("b:%v len(t):%v cap(t):%v \n", b, len(b), cap(b))

	ints := make([]int, 2, 10)
	fmt.Println(len(ints), cap(ints))
}
