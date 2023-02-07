package main

import (
	"bufio"
	f "fmt" // pythonでいうimport fmt as f をやっている感じ
	"os"
	// . "pra_go/animals"  パッケージ名をなくすこともできる
)

func init() {
	s := "このinit関数はmain関数より先に実行される\nそれは初期化用の関数なため"
	f.Println(s)
}

func main() {
scanner := bufio.NewScanner(os.Stdin)
print("何か打ち込んで:::")
var input string
for scanner.Scan() {
	input = scanner.Text()
	break
	}

/*スキャンにエラーが出た時の処理*/

if err := scanner.Err(); err != nil {
	f.Println(os.Stderr, "読み込みエラー", err)
}
	f.Println("標準入力は"+input)
}



func sum(srcs ...int) int {
	//可変長引数はようそがスライスで返される
	num := 0
	for _, src := range srcs {
		num += src
	}
	return num
}

