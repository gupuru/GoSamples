package main

import "fmt"

/**
定数
 */
func main() {

	const P float32 = 2.44444
	fmt.Println(P)

	const a, b, x int = 1, 2, 2
	fmt.Println(a, b, x)

	const (
		ZERO = iota
		ONE = iota
	)

	fmt.Println(ZERO, ONE)

	const (
		THREE = iota
		FOUR
		FIVE
	)

	fmt.Println(THREE, FOUR, FIVE)

}
