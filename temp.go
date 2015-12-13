package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

/**
テンポラリファイルの作成
 */
func main() {

	dir := os.TempDir()

	file, _ := ioutil.TempFile(dir, "temp_test")

	fmt.Println(file.Name())

}