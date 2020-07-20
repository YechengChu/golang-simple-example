package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

// Square has Area() function so can be a Shaper
func (sq *Square) Area() float32 {
	return sq.side * sq.side
} // Area

func main() {
	sq1 := new(Square)
	sq1.side = 3

	var areaIntf Shaper
	areaIntf = sq1
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
} // main
