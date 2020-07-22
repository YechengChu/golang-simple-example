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
		close(naturals)
	}()

	// Squarer
	go func() {
		// 当naturals还有的时候
		// 接收naturals的值到x
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	// 当squares还有的时候
	// 接收squares的值到x
	for x := range squares {
		fmt.Println(x)
	} // for
} // main
