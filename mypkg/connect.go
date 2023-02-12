package mypkg


import (
	//"fmt" // pythonでいうimport fmt as f をやっている感じ
	"fmt"
	"html/template"
	"log"
	"net/http"
	// . "pra_go/animals"  パッケージ名をなくすこともできる
)


func ConnectMain() {
	addition := "user_form"
	addition2 := "user_confirm"
	log.Println("http://localhost:8080/"+addition)
	http.HandleFunc("/"+addition, HandlerUserForm)
	http.HandleFunc("/"+addition2, HandlerUserConfirm)
	http.ListenAndServe(":8080", nil)
}


func HandlerUserForm(w http.ResponseWriter, req *http.Request) {
	/*テンプレートをパースする*/
	tpl := template.Must(template.ParseFiles("./files/templates/test.html"))

	/*テンプレートに出力するマップをセット*/
	values := map[string]string{}//今回は送るべき情報がないからmap内は何もなし
	// マップを展開してテンプレートを出力する
	if err := tpl.ExecuteTemplate(w, "test.html", values);err != nil {
		fmt.Println(err)
	}
}

func HandlerUserConfirm(w http.ResponseWriter, req *http.Request) {
	// テンプレートをパースする
	tpl := template.Must(template.ParseFiles("./files/templates/test_user_confirm.html"))

	// テンプレートに出力する値をマップにセット
	values := map[string]string {
		"account":req.FormValue("account"), 
		"name":req.FormValue("name"), 
		"passwd":req.FormValue("passwd"),
	}

	// マップを展開してテンプレートを出力する
	if err := tpl.ExecuteTemplate(w, "test_user_confirm.html", values); err != nil {
		log.Fatal(err)
	}

}
