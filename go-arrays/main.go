package main

import "fmt"

func main() {

	myarr := [4]int{1, 2, 3, 4}
	myarr[0] = 10
	myarr[3] = 40
	fmt.Println(myarr)

}
func slice() {

	sinf9 := []string{"ali", "vali", "bro"}
	sinf9 = append(sinf9, "hoji")
	fmt.Println(sinf9)
}
