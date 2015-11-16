package main

import "fmt"

type Calculatar interface  {
	Calculatar(a int, b int) int
}

type Add struct  {
}

func (x Add) Calculatar(a int, b int) int {
	return a + b
}

type Sub struct  {
}

func (x Sub) Calculatar(a int, b int) int {
	return a - b
}

/**
インターフェイス
 */
func main() {

	var add Add
	var sub Sub

	var cal Calculatar

	cal = add

	fmt.Println("足し算:", cal.Calculatar(1, 3))

	cal = sub

	fmt.Println("引き算:", cal.Calculatar(3, 2))

}
