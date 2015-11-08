package main

import (
	"fmt"
	"unicode/utf8"
)

/**
文字列
 */
func main() {
	var str string
	str = "はろー"
	fmt.Println(str)

	str = str + " わーるど"
	fmt.Println(str)
	//文字列の長さ
	fmt.Println(str, " 長さ:", len(str))
	//文字数
	fmt.Println(str, " 文字数:", utf8.RuneCountInString(str))

	str = "Hello World"
	fmt.Println(str, " 長さ:", len(str))
	//文字数
	fmt.Println(str, " 文字数:", utf8.RuneCountInString(str))

}