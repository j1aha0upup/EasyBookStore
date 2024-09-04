package model

import (
	"book_store_demo/src/utils"
	"fmt"

	"github.com/google/uuid"
)

type Session struct {
	Uuid      string
	User_name string
	User_id   int
}

func (user *User) Login() (session_temp Session) {
	session_temp.Uuid = uuid.New().String()
	session_temp.User_name = user.Username
	session_temp.User_id = user.Id

	sqlStr := "insert into session value(?,?,?)"

	_, err := utils.Db.Exec(sqlStr, session_temp.Uuid, session_temp.User_name, session_temp.User_id)
	if err != nil {
		fmt.Println("model.session.Login error" + err.Error())
	}

	return
}

func Logout(session_temp string) {
	sqlStr := "DELETE FROM session WHERE uuid=?"
	_, err := utils.Db.Exec(sqlStr, session_temp)
	if err != nil {
		fmt.Println("model.session.Logout error" + err.Error())
	}
}

func CheckLogin(session_uuid string) (session_temp Session) {
	sqlStr := "SELECT uuid,user_name,user_id FROM session WHERE uuid=?"
	result := utils.Db.QueryRow(sqlStr, session_uuid)
	result.Scan(&session_temp.Uuid, &session_temp.User_name, &session_temp.User_id)
	return
}
