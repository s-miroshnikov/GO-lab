// package main

// import "fmt"

// func main() {
// 	i, j := 42, 2701

// 	p := &i         // point to i
// 	fmt.Println(*p) // read i through the pointer
// 	*p = 21         // set i through the pointer
// 	fmt.Println(i)  // see the new value of i

// 	p = &j         // point to j
// 	*p = *p / 35   // divide j through the pointer
// 	fmt.Println(j) // see the new value of j
// }

package main

import "fmt"

func main() {
	a := 1
	b := &a
	fmt.Println("A: ", a)
	fmt.Println("B: ", b)
	fmt.Println("B: ", *b)
}
