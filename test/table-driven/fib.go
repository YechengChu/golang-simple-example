package main

func Fib(n int) int {
	if n < 2 {
		return n
	} // if
	return Fib(n-1) + Fib(n-2)
} // Fib
