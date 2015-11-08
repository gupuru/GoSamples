package main

import "fmt"

/**
for
 */
func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 3 {
			fmt.Println("抜けるよ")
			break
		}
	}

	arr := [...]int{0, 1, 2, 3, 4}

	for i := range arr {
		fmt.Println(i)
	}

}
