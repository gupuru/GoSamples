package main

import (
	"fmt"
	"time"
	"math/rand"
)

const gorotines = 3

/**
チャネル 同期
 */
func main() {

	c := make(chan int)

	for i := 0; i < gorotines; i++ {
		go func(s chan <- int) {
			time.Sleep(time.Duration(rand.Int63n(10)) + time.Second)
			fmt.Println("処理完了")
			s <- 0
		}(c)
	}

	for i := 0; i < gorotines; i++ {
		<-c
	}

	fmt.Println("すべて完了")

}