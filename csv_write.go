package main

import (
	"fmt"
	"os"
	"encoding/csv"
)

/**
csvファイルを出力する
 */
func main() {

	file, err := os.OpenFile("test.csv", os.O_CREATE | os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		file.Close()
	}()

	w := csv.NewWriter(file)

	w.Write([]string{"No.", "商品", "価格", "備考"})
	w.Write([]string{"1", "りんご", "111111", "まあ、いいんちゃう"})
	w.Write([]string{"2", "みかん", "2222", ""})

	w.Flush()

}
