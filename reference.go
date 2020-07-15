package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		// 因为defer了，所以当func执行时i的值为3
		defer func() { println(i) }()
	} // for
	// 因为其他都defer了，所以这行会先执行
	fmt.Println("---------------------------------")
	for i := 0; i < 5; i++ {
		// defer是后进先出，所以最后一个defer会先执行
		// 也就是The value of i is: 2
		defer fmt.Printf("The value of i is: %d\n", i)
	} // for
	fmt.Println("Looks like the end of main func!")
} // main
