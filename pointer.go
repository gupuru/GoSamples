package main

import "fmt"

/**
ポインタ
 */
func main() {

	var ptr *int

	var i int = 1242

	ptr = &i

	fmt.Println("iのアドレス", &i)
	fmt.Println("ptrのアドレス", ptr)

	fmt.Println("iの値", i)
	fmt.Println("ポインタ経由のiの値", *ptr)

	*ptr = 9999999

	fmt.Println("ポインタ経由で変更したiの値", i)
}
