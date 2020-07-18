package main

import "fmt"

type Point struct{ X, Y float64 }

// 这个方法的名字是(*Point).ScaleBy 这里的括号是必须的
func (p *Point) ScaleBy(factor float64) {
  p.X *= factor
  p.Y *= factor
} // ScaleBy

func main() {
  // 调用方法1
  // addP1是一个指向Point{1, 2}的指针，存着Point{1, 2}的地址
  addP1 := &Point{1, 2}
  addP1.ScaleBy(2)
  fmt.Println(*addP1) // "{2, 4}"
  // 调用方法2
  // p2是一个Point
  p2 := Point{3, 4}
  (&p2).ScaleBy(2)
  fmt.Println(p2) // "{6, 8}"
  // 调用方法3
  p3 := Point{5, 6}
  // 编译器会隐式调用 &p3
  p3.ScaleBy(10)
  fmt.Println(p3) // "{50, 60}"
} // main
