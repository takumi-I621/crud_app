package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/createuser", createUser)
	mux.HandleFunc("/deluser", delUser)
	mux.HandleFunc("/upduser", updUser)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

//トップページ
func index(w http.ResponseWriter, r *http.Request) {
	//usersはUserすべてを集めたスライス
	users, err := Users()
	if err != nil {
		log.Fatal("users error")
	}
	//テンプレートの解析
	t := template.Must(template.ParseFiles("form.html"))
	//データに構造たいUsersのスライスを渡す。テンプレートform.htmlを実行
	err = t.ExecuteTemplate(w, "form.html", users)
	if err != nil {
		log.Fatal(err)
	}
}

func Users() {
	panic("unimplemented")
}

//ユーザを追加するハンドラ関数
func createUser(w http.ResponseWriter, r *http.Request) {
	user := Users{}
	//リクエストを解析し、フォームを取得。
	user.Name = r.PostFormValue("name")

	err := user.Create()
	if err != nil {
		log.Fatal("create error")
	}
	http.Redirect(w, r, "/", 302)
}

//ユーザを削除するハンドラ関数
func delUser(w http.ResponseWriter, r *http.Request) {
	//削除する関数
	id, err := strconv.Atoi(r.PostFormValue("id"))
	if err != nil {
		log.Fatal("delのAtoi変換エラー")
		//エラー1：未入力
		//エラー2：削除したいID以外の文字677
		//エラー3：IDの範囲外
	}
	err = Delete(id)
	if err != nil {
		log.Fatal("Deleteエラー")
	}
	http.Redirect(w, r, "/", 302)
}

func Delete(id int) {
	panic("unimplemented")
}

func updUser(w http.ResponseWriter, r *http.Request) {
	user := user{}
	var err error
	//HTMLからフォームの値を取得、構造体に格納する
	user.Id, err = strconv.Atoi(r.PostFormValue("id"))
	if err != nil {
		log.Fatal("upd_Atoi error")
	}
	user.Name = r.PostFormValue("name")
	err = user, Update()
	if err != nil {
		log.Fatal("update error")
	}
	http.Redirect(w, r, "/", 302)
}
