package main

import (
	f "fmt" // pythonでいうimport fmt as f をやっている感じ
	// . "pra_go/animals"  パッケージ名をなくすこともできる
)

/*
	type Person struct {
	    firstName string
	    age int
	}
*/
var s = ""

func init() {
	s = "このinit関数は"
}

func init() {
	s = s + "main関数より"
}

func init() {
	s = s + "先に実行される\nそれは初期化用の関数なため"
}

func main() {
	f.Println(s)
	slices := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "m", "l", "n"}
	slices[0] = "aaaaaaaa"
	var slices2 []string
	for index, slice := range slices {
		f.Printf("[%d] ==> %v\n", index, slice)
		slices2 = append(slices2, slice)
	}
	println("____________________________________")
	f.Println(slices2)
	println("sum関数実行")
	f.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20))
	println("__________________________________")
	println("連想配列作成")
	suwaScore := map[string]int{"国語": 87, "数学": 10, "社会": 98, "英語": 90, "理科": 20}

	f.Println(suwaScore["国語"])
}

func sum(srcs ...int) int {
	num := 0
	for _, src := range srcs {
		num += src
	}
	return num

}

//これがコメント
/*
複数行
*/
