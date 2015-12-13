package main

import (
	"encoding/base64"
	"fmt"
)

/**
base64エンコード, デコード
 */
func main() {

	data := "base64Test"

	//エンコード
	enc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(enc)

	//デコード
	dec, _ := base64.StdEncoding.DecodeString(enc)
	fmt.Println(string(dec))

}