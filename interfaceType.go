package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

type MoreShaper interface {
	Shaper
	Perimeter() float32
}

type Animal interface {
	Bark()
}

type Dog struct{
	weight float32
}

func main() {
	var areaIntf Shaper = &Square{5}
	var aSquare = Square{4}
	var areaIntf2 Shaper = &Circle{2}
	var anAnimal Animal = Dog{21}
	fmt.Println("--------断言类型是一个具体类型--------")
	// Is Square the type of areaIntf ?
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
		fmt.Printf("t is: %v\n", t)
	} // if
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
		fmt.Printf("u is: %v\n", u)
	} // if ... else
	// Is Square the type of &aSquare ?
	if t, ok := (Shaper)(&aSquare).(*Square); ok {
		fmt.Printf("The type of &aSquare is: %T\n", t)
		fmt.Printf("t is: %v\n", t)
	} // if
	fmt.Println("--------断言类型是一个接口--------")
	// is areaIntf a Shaper ?
	if z, ok := areaIntf.(Shaper); ok {
		fmt.Println("The of areaIntf is a Shaper")
		fmt.Printf("z is: %v\n", z)
	} // if
	// is areaIntf2 a MoreShaper ?
	if u, ok := areaIntf2.(MoreShaper); ok {
		fmt.Printf("The type of areaIntf2 is: %T\n", u)
	} else {
		fmt.Println("areaIntf2 is not an interface MoreShaper")
		fmt.Printf("u is: %v\n", u)
	} // if ... else
	// is &aSquare a MoreShaper ?
	if z, ok := (Shaper)(&aSquare).(MoreShaper); ok {
		fmt.Println("The &aSquare is a MoreShaper")
		fmt.Printf("z is: %v\n", z)
	} // if
	fmt.Println("--------switch--------")
	// testing with switch:
	switchType(areaIntf)
	switchType(anAnimal)
	switchInterface((Shaper)(&aSquare))
	switchInterface(anAnimal)
} // main

func switchType(x interface{}) {
	switch t := x.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case Dog:
		fmt.Printf("Type Dog %T with value %v\n", t, t)
	default:
		fmt.Printf("Unexpected type %T", t)
	} // switch
} // switchType

func switchInterface(x interface{}) {
	switch u := x.(type) {
	case MoreShaper:
		fmt.Printf("A MoreShaper type with value %v\n", u)
	case Shaper:
		fmt.Printf("A shaper interface with value %v\n", u)
	default:
		fmt.Printf("Unexpected interface with value %v\n", u)
	} // switch
} // switchInterface

func (sq *Square) Area() float32 {
	return sq.side * sq.side
} // Area

func (d Dog) Bark() {
	switch {
  case d.weight < 10:
		fmt.Println("woof-woof")
	case 10 <= d.weight && d.weight < 20:
		fmt.Println("ruff-ruff")
	default:
		fmt.Println("bow-wow")
	} // switch
} // Area

func (sq *Square) Perimeter() float32 {
	return sq.side * 4
} // Perimeter

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
} // Area
