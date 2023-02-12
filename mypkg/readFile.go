package mypkg

import (
	"fmt"
	"log"
	"os"
)

func ReadFile(filepath string) {
	file, err := os.OpenFile(filepath, os.O_RDONLY | os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 50)
	// Readメソッドでファイルを読み込んで出力
	file.Read(buf)
	fmt.Println(string(buf))

}