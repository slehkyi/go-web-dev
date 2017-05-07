package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

type transportation interface {
	transportationDevice() string
}

func (t truck) transportationDevice() string {
	return fmt.Sprintln("I am a truck with", t.doors, "doors and I am durable.")
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintln("I am a sedan with", s.doors, "doors and I am luxurious.")
}

func report(x transportation) string {
	return fmt.Sprintln(x.transportationDevice())
}

func main() {
	t := truck{
		vehicle{
			4,
			"red",
		},
		true,
	}
	fmt.Println(t)
	s := sedan{
		vehicle{
			2,
			"yellow",
		},
		false,
	}
	fmt.Println(s)
	fmt.Println("The color of the truck is", t.color)
	fmt.Println("The color of the sedan is", s.color)
	fmt.Println(s.transportationDevice())
	fmt.Println(t.transportationDevice())
	fmt.Println(report(s))
	fmt.Println(report(t))
}

/*
Tasks 6-10 from https://goo.gl/LpQDzH
*/
