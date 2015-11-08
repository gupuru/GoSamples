package main

import "fmt"

const C1 = 123456

const C2 int = 33333333333

/**
型なしの計算
 */
func main() {

	var x int = C1 * C2
	fmt.Println(x)

	var a int32 = 12
	var v int64 = 33

	fmt.Println(a + int32(v))
}