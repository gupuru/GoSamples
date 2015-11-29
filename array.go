package main

import "fmt"

/**
配列
 */
func main() {

	var date [3]string

	fmt.Println(len(date))

	date[0] = "日曜日"
	date[1] = "月曜日"
	date[2] = "火曜日"

	for _, value := range date  {
		fmt.Println(value)
	}

	weather := [...]string{"晴れ", "曇り", "雨"}
	fmt.Println(weather)

}
