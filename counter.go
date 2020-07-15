package main

import "fmt"

func getClosure() func() {
	count := 0
	fmt.Printf("The count in getClosure() is %d\n", count)
	return func() {
		count++
		fmt.Println(count)
	}
} // getClosure

func main() {
	counter := getClosure()
	counter()
	counter()
	counter()
	counter()
	counter()
	counter()
	counter()
	counter()
	counter2 := getClosure()
	counter2()
	counter()
} // main
