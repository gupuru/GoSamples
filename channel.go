package main

import "fmt"

/**
チャネル
 */
func main() {

	c := make(chan int)

	go func(s chan<- int) {
		for i := 0; i < 5; i++ {
			s <- i
		}
		close(s)
	}(c)


	for {
		value, ok := <-c

		if !ok {
			break
		}

		fmt.Println(value)
	}

}