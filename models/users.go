package models

import "database/sql"

_ "github.com/mattn/go-sqlite3"

type User struct {
	id int `json:"id"`
	firstName string `json:"firstName"`
	otherNames string `json:"otherNames"`
	email string `json:"email"`
	userName string `json:"userName"`
	password string `json:"password"`
}

type UserCollectionDetails struct {
	Users[]User `json:"users"`
}

func CreateUser(db *sql.DB, name string) (int64 error) {
	sql := "INSERT INTO users(firstName, otherNames, email, userName, password) VALUES(?)"

	Stmt, err1 := db.Prepare(sql)
	if err1 != nil {
		panic(err1)
	}
	defer.Stmt.Close()

	result, err2 := stmt.Exec(firstName, otherNames, email, userName, password)

	if err2 {
		panic(err2)
	}

	return result.LastInsertId()
}
