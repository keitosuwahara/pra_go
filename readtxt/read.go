package readtxt

import (
	"fmt"
	"os"
	"io/ioutil"
)

func Read() {
	file, err := os.Open("../files/test.txt")
	if err != nil {
		fmt.Println("ファイルが読み込めませんでした。")
		os.Exit(1)
	}
	data, err := ioutil.ReadAll(file)
	fmt.Println(string(data))
}