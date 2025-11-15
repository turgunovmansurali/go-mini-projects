package main

import "fmt"

func main() {

	// ozgaruvchilar()
	// if_else()
	// loop1()
	// array()
	// salom_bot()
	// hisobchi()
	// royhat()
	taqqoslash2()
}
func ozgaruvchilar() {
	// var name string = "Ali"
	// var age = 19
	// var height float64 = 1.79

	// fmt.Println("ism: ", name)
	// fmt.Println("yosh: ", age)
	// fmt.Println("bo'y: ", height)

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

func array() {

	myarr := [4]int{1, 2, 3, 4}
	myarr[0] = 10
	myarr[3] = 40
	fmt.Println(myarr)

	sinf9 := []string{"ali", "vali", "bro"}
	sinf9 = append(sinf9, "hoji")
	fmt.Println(sinf9)

}

func if_else() {

	age := 20

	if age < 18 {
		fmt.Println("Siz voyaga yetmagansiz!")
	} else if age == 18 {
		fmt.Println("Siz voyaga yetgansiz!")
	} else {
		fmt.Println("siz ham voyaga yetgansiz :)")
	}

}

func salom_bot() {
	var name string

	fmt.Println("ismingizni kiriting:")
	fmt.Scanln(&name)
	fmt.Println("salom", name, "!")

}
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

func loop1() {

	var a string
	var b string
	var c string

	fmt.Println("1-nima olamiz?")
	fmt.Scanln(&a)
	fmt.Println("2-nima olamiz?")
	fmt.Scanln(&b)
	fmt.Println("3-nima olamiz?")
	fmt.Scanln(&c)

	myslice := []string{}

	myslice = append(myslice, a)
	myslice = append(myslice, b)
	myslice = append(myslice, c)

	fmt.Println()
	fmt.Println("mana toliq ro'yhat!")
	// fmt.Println(myslice[0])
	// fmt.Println(myslice[1])
	// fmt.Println(myslice[2])

	for _, v := range myslice {
		fmt.Println(v)
	}

}

func royhat() {

	myslice1 := []string{}
	var a string

	fmt.Println("Bizga nima kk?")
	fmt.Scanln(&a)

	for {

		fmt.Println("yana nima kk?")
		fmt.Scanln(&a)

		if a != "stop" {

			myslice1 = append(myslice1, a)

		} else {
			fmt.Println()
			fmt.Println()
			fmt.Println("Mana to'liq ro'yhatingiz!")
			fmt.Println()

			for _, v := range myslice1 {
				fmt.Println(v)
			}
			fmt.Println()
			fmt.Println("Bu barcha mahsulotlar soni: ", len(myslice1))
			break
		}
	}
}
func taqqoslash() {

	numbers := []int{1, 50, 6, 2, 20, 50, 11, 3, 4, 8, 10}
	a := len(numbers)
	b := a - 1
	var x int
	var y int

	for i, v := range numbers {

		if i == 0 {
			x = v
		}

		if i == b {
			y = v
		}
	}

	if x > 0 && y < 10 {
		fmt.Println("rost")
	} else {
		fmt.Println("yolg'on")
	}

}
func taqqoslash1() {
	numbers := []int{1, 50, 6, 2, 20, 50, 11, 3, 4, 8, 10, 9, 9, 10, 20}
	// numbers := []int{}
	a := numbers[0]

	b := len(numbers) - 1
	c := numbers[b]

	if a > 0 && c < 10 {
		fmt.Println("rost")
	} else {
		fmt.Println("yolg'on")
	}

}
func taqqoslash2() {
	numbers := []int{1, 50, 6, 2, 20, 50, 11, 3, 4, 8, 10, 9, 9, 10, 20}
	// numbers := []int{} // test uchun

	// 1) Avval slice bo‘shmi, shuni tekshiramiz
	if len(numbers) == 0 {
		fmt.Println("slice bo'sh!")
		return
	}

	// 2) Birinchi element
	a := numbers[0]

	// 3) Oxirgi element
	c := numbers[len(numbers)-1]

	// 4) Taqqoslash
	if a > 0 && c < 10 {
		fmt.Println("rost")
	} else {
		fmt.Println("yolg'on")
	}
}
