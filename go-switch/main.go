package main

import "fmt"

func main() {

	a := 5
	switch a {
	case 5:
		fmt.Println("yaxshi")
	case 4:
		fmt.Println("qoniqarli")
	case 3:
		fmt.Println("qoniqarsiz")
	default:
		fmt.Println("Yomon!")
	}

	//2-usuli bu usulda, bir vaqtda 2ta shart bajarish mumkin

	color := "orange"

	switch color {
	case "red", "orange":
		fmt.Println("Stop")
	case "green", "black":
		fmt.Println("Go")
	case "yellow":
		fmt.Println("Wait")
	default:
		fmt.Println("Unknow color")
	}

	b := 100
	switch {
	case b < 95 || b > 85:
		fmt.Println("yaxshi")
	case b < 84 || b > 65:
		fmt.Println("qoniqarli")
	case b < 64 && b > 55:
		fmt.Println("qoniqarsiz")
	default:
		fmt.Println("Yomon!")
	}

}
