package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 20; x++ {
			// 将x的值传入通道naturals
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			// 从naturals通道中拿出一个int，并将其赋予x
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for x := 0; x < 10; x++ {
		// 从squares通道中拿出一个int打印出来
		fmt.Println(<-squares)
	} // for
} // main
