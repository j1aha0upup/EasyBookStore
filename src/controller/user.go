package controller

import (
	"book_store_demo/src/model"
	"html/template"
	"net/http"
)

func Login_verify(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	user := model.Login_verify(username, password)
	session_gettemp, _ := r.Cookie("user_session")
	session_temp := session_gettemp.String()
	if session_temp == "" {
		if user.Id != 0 {
			session := user.Login()
			cookie := http.Cookie{
				Name:   "user_session",
				Value:  session.Uuid,
				MaxAge: 60,
			}
			http.SetCookie(w, &cookie)
			cookie = http.Cookie{
				Name:   "user_name",
				Value:  session.User_name,
				MaxAge: 60,
			}
			http.SetCookie(w, &cookie)

			template.Must(template.ParseFiles("src/pages/user/login_success.html")).Execute(w, struct {
				User_name string
			}{
				User_name: user.Username,
			})

		} else {
			template.Must(template.ParseFiles("src/pages/user/login.html")).Execute(w, struct {
				ClueStr  string
				username string
				password string
			}{
				ClueStr:  "用户名 or 密码 fault",
				username: username,
				password: password,
			})
		}
	} else {
		template.Must(template.ParseFiles("src/pages/user/login_success.html")).Execute(w, struct {
			User_name string
		}{
			User_name: user.Username,
		})
	}
}

func Regist_verify(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	user := model.CheckUsername(username)

	if user.Id == 0 {
		user.Username = username
		user.Password = r.PostFormValue("password")
		user.Email = r.PostFormValue("email")
		err := user.AddUser()

		if err != nil {
			panic(err)
		}
		template.Must(template.ParseFiles("src/pages/user/regist_success.html")).Execute(w, nil)
	} else {
		template.Must(template.ParseFiles("src/pages/user/regist.html")).Execute(w, struct {
			ClueStr  string
			username string
			password string
			email    string
		}{
			ClueStr:  "uesrname already exists",
			username: username,
			password: user.Password,
			email:    user.Email,
		})
	}
}
