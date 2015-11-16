package main

import "fmt"

/**
アサーション
 */
func main() {

	var i interface{} = "test"

	s1, ok := i.(string)
	fmt.Println(s1, ok)
	if ok {
		fmt.Println(s1 + "だよ")
	}

	s2, ok := i.(interface{
		dummy()
	})

	fmt.Println(s2, ok)

}
