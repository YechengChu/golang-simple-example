package main

func main() {
	var g int
	for i := 0; i < 3; i++ {
		func(s string) {
			g = i
			println(g)
			println(s)
		}("Hello World!")
	} // for
} // main
