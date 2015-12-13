package main

import (
	"fmt"
	"os"
	"encoding/csv"
)

/**
csvファイルを読み込む
 */
func main() {

	file, err := os.OpenFile("test.csv", os.O_RDONLY, 0)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		file.Close()
	}()

	r := csv.NewReader(file)

	for  {
		record, err := r.Read()
		if err != nil {
			break
		}

		fmt.Println(record)
	}
}
