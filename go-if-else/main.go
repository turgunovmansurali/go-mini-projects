package main

import "fmt"

func main() {

	age := 20 //16-18-20 yosh qoysa shartlar tekshiriladi

	if age < 18 {
		fmt.Println("Siz voyaga yetmagansiz!")
	} else if age == 18 {
		fmt.Println("Siz voyaga yetgansiz!")
	} else {
		fmt.Println("siz ham voyaga yetgansiz :)")
	}

}
