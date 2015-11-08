package main

import "fmt"

/**
関数
 */
func main() {
	fmt.Println(plus(1, 5))

	isBig := big(2, 3)
	fmt.Println(isBig)

	add, sub, mul, div := calc(5, 3)
	fmt.Println(add, sub, mul, div)

	f2(f1())

	holiday(1, "元旦", "成人の日")
	holiday(3, "春分の日")

	fmt.Println(calc2(2 ,4))

	val := 22222

	func(i int) {
		fmt.Println(i * val)
	}(11)

	f := func(s string) {
		fmt.Println(s)
	}

	f("クロージャ")
}

/**
足し算
 */
func plus(a int, b int) int {
	return a + b
}

/**
大小
 */
func big(a int, b int) bool {
	if a > b {
		return true
	} else {
		return false
	}
}

/**
多値
 */
func calc(a int, b int) (int, int, int, float32) {
	return a + b, a - b, a * b, float32(a) / float32(b)
}

/**
多値(返り値に名前あり)
 */
func calc2(a int, b int) (add int, sub int, mul int, div float32) {
	add = a + b
	sub = a - b
	mul = a * b
	div = float32(a) / float32(b)

	return
}

func f1() (int, string, float32) {
	return 0, "niku", 2.2222
}

func f2(a int, b string, c float32) {
	fmt.Println(a, b, c)
}

/**
可変長
 */
func holiday(month int, days ...string) {

	fmt.Printf("%d月の祝日は%d日あるよ", month, len(days))
	for _, day := range days {
		fmt.Println(day)
	}

}