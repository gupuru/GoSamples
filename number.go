package main

import "fmt"

/**
数値
 */
func main() {
	//普通のint
	var i int = 123456789

	fmt.Println(i)

	//int 64の場合
	var i64 int64 = int64(i)

	fmt.Println(i64)
}