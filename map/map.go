package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	if m == nil {
		fmt.Println("m is nil")
	}
	m = make(map[string]Vertex)
	m["Benaulim"] = Vertex{
		15.257080, 73.918886,
	}
	fmt.Println(m["Benaulim"])

	m[""] = Vertex{5, 5}
	x, ok := m["bjhgjhg"]

	fmt.Println(x, ok)
}
