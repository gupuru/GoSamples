package main

import (
	"fmt"
	"time"
)

/**
ゴルーチン
 */
func main() {
	
	fmt.Println("main start")

	fmt.Println("default test start")
	test()

	fmt.Println("go test start")
	go test()

	time.Sleep(3 * time.Second)

	fmt.Println("main finish")
}

func test()  {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}