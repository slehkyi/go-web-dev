package main

import "fmt"

type person struct {
	fName string
	lName string
	favFood []string
}

func (p person) walk() string {
	return fmt.Sprintln(p.fName, "is walking")
}

func main() {
	p1 := person{"Serhii", "Lehkyi", []string{"meat", "eggs"}}
	fmt.Println(p1)
	fmt.Println(p1.fName)
	fmt.Println(p1.favFood)
	for _, v := range  p1.favFood {
		fmt.Println(v)
	}

	s := p1.walk()
	fmt.Println(s)
}

/*
Tasks 1-5 from https://goo.gl/LpQDzH
 */
