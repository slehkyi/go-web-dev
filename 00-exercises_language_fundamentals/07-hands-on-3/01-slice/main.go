package main

import "fmt"

func main() {
	sliceInt := []int{1, 3, 4, 5}
	fmt.Println(sliceInt)
	for i, _ := range sliceInt {
		fmt.Println(i)
	}
	for i, v := range sliceInt {
		fmt.Println(i, "-", v)
	}
}
