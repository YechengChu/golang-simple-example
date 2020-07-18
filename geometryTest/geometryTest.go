package main

import (
	"fmt"

	"./geometry"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}
	fmt.Println(geometry.Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))           // "5", method call
	fmt.Println("---------------------------------------------")
	perim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	// 下行代码不能通过编译报错信息为
	// not enough arguments in call to geometry.Distance
	//   have (geometry.Path)
	//   want (geometry.Point, geometry.Point)
	// fmt.Println(geometry.Distance(perim)) // "12", standalone function
	// 如果要让上面代码可以执行，可将geometry包中的Distance名字改为PathDistance，再调用
	// fmt.Println(geometry.PathDistance(perim))
	fmt.Println(perim.Distance()) // "12", method of geometry.Path
}
