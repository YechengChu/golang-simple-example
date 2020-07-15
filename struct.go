package main

import "fmt"

type animal struct {
	name       string
	animalType string
	weight     int
	age        int
}

func addWeight1(a animal, w int) animal {
	a.weight += w
	return a
} // addWeight1

// 使用指针传入和返回结构体提高效率
func addWeight2(a *animal, w int) *animal {
	a.weight += w
	return a
} // addWeight2

func main() {
	// animal1的类型是animal
	var animal1 animal = animal{"lion", "feline", 230, 2}
	fmt.Println(animal1)
	fmt.Printf("The type of animal1 is: %T\n", animal1)
	fmt.Printf("animal1 has name: %s\ttype:%s\tweight:%d\tage:%d\n", animal1.name, animal1.animalType, animal1.weight, animal1.age)
	// animal2的类型是*animal，其实&Type{}相当于调用new(T)
	var animal2 = &animal{age: 3, name: "tiger", weight: 12}
	fmt.Println(animal2)
	fmt.Printf("The type of anima2 is: %T\n", animal2)
	// *aniaml2的*可以省略
	fmt.Printf("animal2 has name: %s\ttype:%s\tweight:%d\tage:%d\n", animal2.name, (*animal2).animalType, animal2.weight, animal2.age)
	// animal3的类型是*animal
	animal3 := new(animal)
	animal3.animalType = "canine"
	animal3.age = 3
	// *可以省略，但如果要加*号，括号必须要加
	(*animal3).name = "wolf"
	fmt.Println(animal3)
	fmt.Printf("The type of animal3 is: %T\n", animal3)
	fmt.Printf("animal3 has name: %s\ttype:%s\tweight:%d\tage:%d\n", animal3.name, animal3.animalType, (*animal3).weight, animal3.age)
	fmt.Println("-------------------------------------------------------------------------------------")
	animal1 = addWeight1(animal1, 12)
	animal1.age = 0
	fmt.Printf("animal1 has name: %s\ttype:%s\tweight:%d\tage:%d\n", animal1.name, animal1.animalType, animal1.weight, animal1.age)
	animal1New := addWeight2(&animal1, 10)
	animal1New.age = 1
	fmt.Printf("animal1 has name: %s\ttype:%s\tweight:%d\tage:%d\n", animal1New.name, animal1New.animalType, animal1New.weight, animal1New.age)
} // mainl
