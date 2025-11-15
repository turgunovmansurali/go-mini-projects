package main

import "fmt"

func main() {
	var name string = "Ali"
	var age = 19
	var height float64 = 1.79

	fmt.Println("ism: ", name)
	fmt.Println("yosh: ", age)
	fmt.Println("bo'y: ", height)

	ozgaruvchilar()
}
func ozgaruvchilar() {
	var x int = 5
	var y int = x + 3
	fmt.Println(y)

	var age float32 = 19
	var height float32 = 1.78
	total := age + height

	fmt.Println(total)

	const PI = 3.14
	a := 1.86
	b := PI + a

	fmt.Println(b)
}
