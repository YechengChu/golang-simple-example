package main

import "fmt"

type people struct {
	name    string
	age     int
	job     string
	address string
}

func greeting(mode int) (msg string) {
	switch mode {
	case 1:
		msg = "Good morning"
	case 2:
		msg = "Hello world!"
	case 3:
		msg = "Hi, nice to meet you!"
	default:
		msg = "Hi!"
	} // switch
	return
} // greeting

func createStruct() *people {
	var p *people = new(people)
	*p = people{"Darren", 20, "Student", "Suzhou"}
	return p
} // createStruct

func greeting2(g func(int) string) {
	fmt.Println("-------------------------------------")
	for i := 0; i <= 5; i++ {
		fmt.Println(g(i))
	} // for
	fmt.Println("-------------------------------------")
} // greeting 2

func main() {
	var g func(int) string
	if g != nil {
		fmt.Println(g(2))
	} else {
		g = greeting
		fmt.Println(g(3))
	}
	greeting2(g)
	c := createStruct
	fmt.Printf("%T\n", c)
	fmt.Println(c)
	var p1 *people = c()
	p1.name = "Tim"
	(*p1).age = 16
	fmt.Println(p1)
	fmt.Println(*c())
} // main
