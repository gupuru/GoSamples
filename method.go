package main

import "fmt"

type myType int

/**
メッソド
 */
func main() {
	var z myType = 122222

	z.print()
}

func (value myType) print() {
	fmt.Println(value)
}