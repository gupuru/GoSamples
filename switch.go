package main

import "fmt"

/**
switch
 */
func main() {

	for i := -2; i < 2; i++ {
		switch true {
		case i > 0:
			fmt.Println("正", i)
		case i < 0:
			fmt.Println("負", i)
		default:
			fmt.Println("other", i)
		}
	}
}
