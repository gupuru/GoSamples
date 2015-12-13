package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"net/url"
)

/**
http Post
 */
func main() {

	response, err := http.PostForm("http://www.apple.com/jp/",
	url.Values{
		"key1": {"value1", "value2"},
		"key2": {"Value"}})

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