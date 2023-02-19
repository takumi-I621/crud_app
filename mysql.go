package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql,", "tm:tm55@/tm")
	if err != nil {
		panic(err)
	}
}

//----------------CRUD処理を行う関数--------------------
func (user *User) Create() (err error) {
	_, err = Db.Exec("INSERT INTO tm (name) VALUES (?)", user.Name)
	if err != nil {
		return
	}
	//SQLを実行、1行だけ値を返す。
	err = Db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&user.Id)
	return
}

func Delete(id int) (err error) {
	_, err = Db.Exec("DELETE FROM tm WHERE id =?", id)
	return
}

//全てのユーザを照会する
func Users() (users []User, err error) {
	rows, err := Db.Query("SELECT * FROM tm")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := User{}
		if err = rows.Scan(&conv.Id, &conv.Name); err != nil {
			return
		}
		users = append(users, conv)
	}
	rows.Close()
	return
}

func (user *User) Update() (err error) {
	Db.Exec("UPDATE tm SET name = ? WHERE id = ?", user.Name, user.Id)
	return
}
