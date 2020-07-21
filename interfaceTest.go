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

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
} // Area

func main() {
	sq1 := new(Square)
	sq1.side = 3.5

	var areaIntf Shaper
	areaIntf = sq1
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Println("The square1 has area: ", areaIntf.Area())

	fmt.Println("---------------------------------------")

	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	sq2 := &Square{5}    // Area() of Square needs a pointer
	// The above two lines the same effect as the following 6 lines
	// rAddr := new(Rectangle)
	// rAddr.length = 5
	// rAddr.width = 3
	// r := *rAddr
	// sq2 := new(Square)
	// sq2.side = 5

	// shapes := []Shaper{Shaper(r), Shaper(sq2)}
	// or shorter
	shapes := []Shaper{r, sq2}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	} // for
} // main
