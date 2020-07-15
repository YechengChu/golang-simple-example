package main

import (
	"fmt"
	"log"
)

func protect(g func()) {
	defer func() {
		log.Println("done protect")
		// Println executes normally even if there is a panic
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	log.Println("start protect")
	g() // have runtime-error
	fmt.Println("I am doing something in protect function!")
} // prorect

func canNotProtect(g func()) {
	defer recover()
	log.Println("start can not protect function")
	g() // have runtime-error
	fmt.Println("I am doing something in protect function!")
} // prorect

func raisePanic() {
	panic("I am very panic!!! :-(")
	fmt.Println("I am doing something in raisePanic function!")
} // raisePanic

func main() {
	panicFunc := raisePanic
	protect(panicFunc)
	fmt.Println("I am doing something in main function!")
	canNotProtect(panicFunc)
} // main
