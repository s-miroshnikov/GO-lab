package main

import "fmt"

func main() {
	x := make([]int, 0, 2)
	x = append(x, 4)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))

}
