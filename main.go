package main

import "fmt"

func hisobchi() {

	var x int
	var y int
	fmt.Println("salom bu dastur 2 sonni o'rta arifmetigini topib beradi!")
	fmt.Println("1-sonni kriting:")
	fmt.Scanln(&x)
	fmt.Println("2-sonni kiriting:")
	fmt.Scanln(&y)

	total := x + y
	z := total / 2

	fmt.Println("natija:", z)

}