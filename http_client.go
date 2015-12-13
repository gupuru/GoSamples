package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
)

/**
http GET
 */
func main() {

	response, err := http.Get("http://www.apple.com/jp/")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("status: ", response.Status)

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(body))

}