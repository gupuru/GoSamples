package main

import "fmt"

/**
型switch
 */
func main() {

	showType(nil)
	showType(2)
	showType("ssss")

}

func showType(x interface{})  {
	switch x.(type) {
	case nil:
		fmt.Println("nil")
	case int, int32:
		fmt.Println("整数")
	default:
		fmt.Println("不明")
	}
}