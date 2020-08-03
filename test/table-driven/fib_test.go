package main

import "testing"

// 对号（√）
const checkMark = "\u2713"
// 叉号（×）
const ballotX = "\u2717"

func TestFib(t *testing.T) {
	var fibTests = []struct {
		in       int // input value
		expected int // expected result
	}{
		// pairs of input value and expected result
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		// add a failing test
		// {7,10},
		{7, 13},
	}

	// for elements in fibTests
	for _, tt := range fibTests {
		// call Fib function use in value
		actual := Fib(tt.in)
		if actual != tt.expected {
			// if actual result not equal to expected result
			// print error
			t.Errorf("%v Fib(%d) = %d; expected %d", ballotX, tt.in, actual, tt.expected)
		} else {
			// else, print log
			t.Logf("%v Fib(%d) = %d", checkMark, tt.in, actual)
		}
	} // for
} // TestFib

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }

func benchmarkFib(i int, b *testing.B) {
    for n := 0; n < b.N; n++ {
        Fib(i)
    } // for
} // benchmarkFib
