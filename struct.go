package main

import "fmt"

type Person struct {
	name string
	age  int
}

/**
構造体
 */
func main() {

	var x struct {
		a int
		b bool
		c string
	}

	x.a = 1
	x.b = true
	x.c = "aaaa"

	fmt.Println(x)

	p1 := Person{age:11, name:"nixk"}
	fmt.Println(p1)
	p2 := Person{"aaa", 2222}
	fmt.Println(p2)

}
