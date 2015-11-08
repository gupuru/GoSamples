package main

import "fmt"

/**
ラベル
 */
func main() {

	LBL:

	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)

		for j := 0; j < 3; j++ {
			fmt.Println("        j:", j)

			if i == 1 && j == 1 {
				continue LBL
			}
		}

	}

}
