package main

import (
	"github.com/franela/goreq"
	"fmt"
	"os"
)

/**
goreqのテスト
 */
func main() {

	res, err := goreq.Request{ Uri: "http://www.google.com" }.Do()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res.Body.ToString())

}