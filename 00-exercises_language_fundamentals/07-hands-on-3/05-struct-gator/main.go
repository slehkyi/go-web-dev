package main

import "fmt"

type gator int

type flamingo bool

type swampCreature interface {
	greeting()
}

func (g gator) greeting() {
	fmt.Println("Hello, I am gator")
}

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful.")
}

func bayou(sc swampCreature) {
	sc.greeting()
}

func main() {
	var g1 gator
	g1 = 42
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = int(g1)
	fmt.Println(x)

	g1.greeting()
	bayou(g1)
	var f flamingo
	bayou(f)
}

/*
Tasks 11-16 from https://goo.gl/LpQDzH
*/
