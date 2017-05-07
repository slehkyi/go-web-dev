package main

import (
	"fmt"
)

func main() {
	// slice
	// composite literal; slice literal
	x := []int{7, 9, 42}
	fmt.Println(x)

	y := make([]int, 0, 100)
	y = append(y, 8)
	y = append(y, 12)
	y = append(y, 43)
	fmt.Println(y)

	// map
	m := map[string]string{
		"index": "value",
	}
	fmt.Println(m)

	m1 := make(map[int]int)
	m1[1] = 1
	fmt.Println(m1)

	// struct
	type str struct {
		field1 string
		field2 int
	}
	fmt.Println(str{"dtring", 1})
}
