package main

func main() {
	for i := 0; i < 3; i++ {
		func() {
			println(i)
		}
	} // for
} // main
