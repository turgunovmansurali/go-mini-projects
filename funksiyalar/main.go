package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("istalgan raqamni kiriting:")
	fmt.Scanln(&a)
	fmt.Print("qo'shmoqchi bolgan raqamni kiriting:")
	fmt.Scanln(&b)
	fmt.Println("Natija:", hisobchi(a, b))

}

func hisobchi(a int, b int) int {
	var c = a + b
	return c
}
