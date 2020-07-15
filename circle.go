package main

import (
	"fmt"
	"os"
	"strconv"
)

const PI = 3.14

func main() {
	var radiusString = os.Args[1]
	var radius, err = strconv.ParseFloat(radiusString, 32)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error message: %v\n", err)
		os.Exit(1)
	} // if
	var area = calculateArea(radius)
	fmt.Printf("The circle with radius %.2f has an area of %.2f\n", radius, area)
} // main

func calculateArea(r float64) float64 {
	var area = PI * r * r
	return area
} // calculateArea
