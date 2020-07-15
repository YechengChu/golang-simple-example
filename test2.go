package main

import "fmt"

func main() {
	// arr1是一个指向一个5个int的指针
	var arr1 = new([5]int)
	arr2 := *arr1
	arr3 := arr1
	arr2[0] = 18
	arr2[2] = 20
	fmt.Printf("arr1 is: %T\n", arr1)
	fmt.Printf("*arr1 is: %T\n", *arr1)
	fmt.Printf("arr2 is: %T\n", arr2)
	fmt.Println("------------------------------------")
	fmt.Println(arr1)
	fmt.Println(*arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println("------------------------------------")
	fmt.Println("arr1")
	for _, x := range arr1 {
		fmt.Printf("%d\n", x)
	} // for
	fmt.Println("------------------------------------")
	fmt.Println("arr2")
	for _, x := range arr2 {
		fmt.Printf("%d\n", x)
	} // for
	fmt.Println("------------------------------------")
	// change the third element in arr1 to 8
	arr1[2] = 8
	(*arr1)[3] = 9999
	// and the last element in arr3 to 78
	arr3[4] = 78
	fmt.Println("arr1")
	for _, x := range arr1 {
		fmt.Printf("%d\n", x)
	} // for
	fmt.Println("------------------------------------")
	fmt.Println("arr2")
	for _, x := range arr2 {
		fmt.Printf("%d\n", x)
	} // for
	fmt.Println("------------------------------------")
	fmt.Println("arr3")
	for _, x := range arr3 {
		fmt.Printf("%d\n", x)
	} // for
} // main
