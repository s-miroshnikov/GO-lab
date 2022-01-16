package main

import "fmt"

type Creature struct {
	Name     string
	Greeting string
}

func (f Creature) Greet() {
	fmt.Printf("%s says %s", f.Name, f.Greeting)
}

func main() {
	sammy := Creature{
		Name:     "Sammy",
		Greeting: "Hello!",
	}
	Creature.Greet(sammy)
}
