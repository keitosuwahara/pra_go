package main

import (
    "fmt"
    "pra_go/animals"
    "strconv"
    "pra_go/readtxt"
)

/*
type Person struct {
    firstName string
    age int
}
*/

func main() {
    defer fmt.Printf("%v",1)
    animals.Dogfeed()
    ar := [...] string{"aa", "dd","ffd"}
    fmt.Println(ar)

    //スライス
    slice := [] string{"スライスと言ってるが", "リストみたいなもん"}
    var newsli = append(slice, "なんだよなあ")
    lennewsli := strconv.Itoa(len(newsli))
    newsli[2] = "ですよ"
    fmt.Println(newsli, "要素数は"+ lennewsli)

    //連想配列
    var score = map[string]int{"国語":89,"英語":43,"社会":100}
    score["国語"] = 100
    kokugo := strconv.Itoa(score["国語"])
    fmt.Println(kokugo+"点")//delete(score, "国語")でscoreの国語を削除できる

    //for文rangeはpythonのenumerateみたいなやつ
    langs := [] string{"go", "java","python","c","c++","javascript"}
    for index, lang := range langs {
        fmt.Println(index, lang)
    }
    //連想配列for順序はバラバラになってしまうことがある
    for key, value := range score {
        fmt.Println(key, value)
    }

    //indexか要素の変数どっちかしか使わない時 _は使わなくてもいいようだ
    for index := range langs {
        fmt.Println(index)
    }

    //continueを使ってみる
    for _, lang := range langs {
        if lang == "java" {
            fmt.Println(lang, "is java")
            continue

        }
        if lang == "c" {
            fmt.Println(lang, "favorite lang")
            break
        }
        fmt.Println(lang)

        readtxt.Readmain("./files/test.txt")
    }

    /*
    switch文
    var  lang = "jav"

    switch lang {
    case "java":
        print("javaだよ")
    case "python":
        print("pythonだよ")
    case "Go":
        print("Goだよ")
    default:
        println("defaultはelseみたいなもん")
    }
    var ar[3] string
    fmt.Scan(&ar[0], &ar[1], &ar[2])
    fmt.Println(ar)
    */

}




//これがコメント
/*
複数行*/