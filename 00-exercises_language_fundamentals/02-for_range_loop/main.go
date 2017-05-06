package main

import (
"fmt"
)

func main() {
	// slice
	// composite literal; slice literal
	x := []int{7, 9, 42}
	for i , _ := range x {
		fmt.Println(i, "-", x[i])
	}

	y := make([]int, 0, 10)
	y = append(y, 777)

	for i, v := range y {
		fmt.Println(i, "-", v)
	}

	// map
	// composite literal; map literal
	m := map[string]int{"Todd": 45, "Nina": 25, "Patrick": 27}
	for k, v := range m {
		fmt.Println(k, "-", v)
	}

	// slice
	// composite literal; map literal
	m1 := map[string]int{"Todd": 45, "Nina": 25, "Patrick": 27}
	for k, _ := range m1 {
		fmt.Println(k, "-", m1[k])
	}

	// struct
}
