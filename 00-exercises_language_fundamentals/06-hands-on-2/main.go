package main

import "fmt"

type person struct {
	fname string
	lname string
	age int
}

type secretAgent struct {
	person
	licenseToHaveWeapon bool
}

func (p person) speak() {
	fmt.Println("Hey, my name is", p.fname)
}

func (sa secretAgent) speak() {
	fmt.Println("Good morning, kids. I am", sa.lname, ",", sa.fname, sa.lname)
}

type human interface {
	speak()
}

func greeting(g human) {
	g.speak()
}

func main() {
	p1 := person{"Serhii", "Lehkyi", 23}
	sa1 := secretAgent{person{
		"James",
		"Bond",
		40,
	},
	true}
	fmt.Println(p1.fname)
	p1.speak()
	fmt.Println(sa1.fname, sa1.lname)
	sa1.speak()
	greeting(p1)
	greeting(sa1)
}