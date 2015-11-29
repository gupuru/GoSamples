package main

import "fmt"

/**
スライス
 */
func main() {

	x := [3]string{"a", "n", "d"}
	var s1 []string
	s1 = x[:]
	fmt.Println(s1)
	fmt.Println("len:", len(x))
	fmt.Println("cap:", cap(s1))

	data := []int{100, 20, 20}
	fmt.Println(data)
	data = append(data, 100)
	fmt.Println(data)

	s2 := make([]int, 10, 20)
	fmt.Println(s2)
	fmt.Println("len", len(s2))
	fmt.Println("cap", cap(s2))

}
