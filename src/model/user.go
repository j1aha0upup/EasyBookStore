package model

import (
	"book_store_demo/src/utils"
)

type User struct {
	Id       int
	Username string
	Password string
	Email    string
}

func (user *User) AddUser() error {
	sqlStr := "insert into user(username,password,email) values(?,?,?)"
	_, err := utils.Db.Exec(sqlStr, user.Username, user.Password, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func CheckUsername(username string) *User {
	user := &User{}
	sqlStr := "select id,username,password,email from user where username = ?"
	queryResult := utils.Db.QueryRow(sqlStr, username)
	queryResult.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	return user
}

func Login_verify(username string, password string) *User {
	user := &User{}
	sqlStr := "select id,username,password,email from user where username = ? and password = ?"
	queryResult := utils.Db.QueryRow(sqlStr, username, password)
	queryResult.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	return user
}

func CheckAdmin(user_id int) bool {
	temp := 0
	sqlStr := "select user_role from user where id = ?"
	row := utils.Db.QueryRow(sqlStr, user_id)
	row.Scan(&temp)

	return temp == 3
}
