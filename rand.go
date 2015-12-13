package main

import (
	"math/rand"
	"time"
	"fmt"
)

/**
base64エンコード, デコード
 */
func main() {

	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Int())

	fmt.Println(rand.Float32())

	fmt.Println(rand.Float64())

}