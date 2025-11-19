package main

import "fmt"

func main() {

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

//for loop boyicha 2-loyiha
//bunda istalgancha narsalarni royhatini yaratish mumkin!
