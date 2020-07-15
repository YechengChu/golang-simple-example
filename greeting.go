package main

import "fmt"

func greeting(names ...string) {
	for _, i := range names {
		fmt.Printf("Hello %s!\n", i)
	} // for
} // greeting

func main() {
	greeting("Jim", "Sean", "Barry", "Gareth", "Tobby")
} // main
