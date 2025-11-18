package main

import "fmt"

func main() {

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

	for _, v := range myslice {
		fmt.Println(v)
	}

}

//loop boyicha 1-mini loyiha
//cheklangan royhat
