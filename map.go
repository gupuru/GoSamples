package main

import "fmt"

/**
map
 */
func main() {

	country := make(map[string]string)
	country["japan"] = "東京"
	country["us"] = "ワシントン"

	fmt.Println(country["japan"])
	fmt.Println(country["us"])

	value, isCountry := country["イギリス"]
	if isCountry {
		fmt.Println(value)
	} else {
		fmt.Println("error")
	}

}
