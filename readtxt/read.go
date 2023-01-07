package readtxt

import (
	"fmt"
	"os"
)

func Readmain(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err, "です")
	}
	//deferで関数の終了時にファイルを閉じる	
	defer file.Close()

	buf1 := make([]byte, 10)
	buf2 := make([]byte, 10)

	file.Read(buf1)
	fmt.Println(string(buf1))

	file.ReadAt(buf2, 10)
	fmt.Println(string(buf2))
}