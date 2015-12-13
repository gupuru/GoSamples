package main

import "os"

const TEST_DIR string = "testDir"

/**
ディレクトリを作成・削除する
 */
func main() {

	mkDir()
	removeDir()

}

func mkDir()  {
	os.Mkdir(TEST_DIR, 0777)
	os.MkdirAll("a/te/c", 0777)
}

func removeDir(){
	os.Remove(TEST_DIR)
}
