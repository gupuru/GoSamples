package main


import (
	"fmt"
	"os"
)

/**
遅延
 */
func main() {
	fmt.Println("start")
	file, err := os.OpenFile("test.txt", os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString("書き込み")
	fmt.Println("end")
}