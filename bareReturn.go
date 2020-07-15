package main

import "fmt"

func HelloWorld() (message string, num int) {
	message = "Hello world!"
	num = 3
	return
} // HelloWorld

func main() {
	msg, no := HelloWorld()
	fmt.Printf("%s\n%d\n", msg, no)
}
