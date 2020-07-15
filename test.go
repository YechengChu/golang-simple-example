package main

import "fmt"

func main() {
	var p1 = new([]int)
	var p2 = make([]int, 0)
	fmt.Printf("p1 has type %T\n", p1)
	fmt.Println(p1)
	fmt.Printf("p2 has type %T\n", p2)
	fmt.Println(p2)
	*p1 = append(*p1, 2)
	p2 = append(p2, 3)
	fmt.Printf("The first element in p1 is %d.\nThat of p2 is %d.\n", (*p1)[0], p2[0])
} // main
