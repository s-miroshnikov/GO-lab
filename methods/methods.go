package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func (Vertex) Hi() {
	fmt.Println("Hello!")
}
func (v *Vertex) AddX() {
	v.X += 10
}
func main() {
	v := Vertex{33, 25}
	fmt.Println(v.Abs())
	v.Hi()
	x := &v
	x.AddX()
	fmt.Println(x)
}
