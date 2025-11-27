package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "mansurali" //istalgan matin (ism)
	age := 20           // istalgan raqam 'int' (yosh) yozish mumkin

	userName, userAge := userInfo(name, age) //funksiyadan qaytgan javoblarni 2 ta o'zgaruchiga saqlab ularni print qildim
	fmt.Println(userName)
	fmt.Println(userAge)

}

//bu funksiya ismni katta harflarda yoib beradi va yoshni 5 ga oshirib beradi

func userInfo(name string, age int) (string, int) {

	var a = strings.ToUpper(name)
	var b = age + 5
	return a, b
}
