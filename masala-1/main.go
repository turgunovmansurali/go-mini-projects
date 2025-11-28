// 1 dan 10 gacha bolgan sonlar yig'indisini topish !
package main

import "fmt"

func main() {

	i := 10 * (10 + 1) / 2
	fmt.Println(i)
	//bu matematik  S = n × (n + 1) / 2 formula asosida yozilgan

	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	//bu esa go for loop orqali topilgan

}
