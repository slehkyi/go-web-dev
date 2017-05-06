package main

import "fmt"

func main() {
	m := map[string]int{
		"first": 1,
		"second": 2,
	}
	fmt.Println(m)
	for i, _ := range m {
		fmt.Println(i)
	}
	for i, v := range m {
		fmt.Println(i, "-", v)
	}
}
