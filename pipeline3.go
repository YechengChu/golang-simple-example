package main

import "fmt"

// 在counter中out只能用于发送int
func counter(out chan<- int) {
	for x := 0; x < 20; x++ {
		out <- x
	} // for
	close(out)
} // counter

// 在squarer中in只能用于接收int，out只能用于发送int
func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	} // for
	close(out)
} // squarer

// 在printer中in只能用于接收int
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	} // for
} // printer

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
} // main
