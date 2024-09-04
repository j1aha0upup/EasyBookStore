package main

import (
	"book_store_demo/src/controller"
	"book_store_demo/src/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/index.html"))
	IndexPage := &model.Pages{}
	IndexPage.EveryPageRecordCount = 4
	IndexPage.CurrentPage, _ = strconv.Atoi(r.FormValue("currentPage"))
	IndexPage.Min, _ = strconv.ParseFloat(r.FormValue("min"), 64)
	IndexPage.Max, _ = strconv.ParseFloat(r.FormValue("max"), 64)
	if IndexPage.CurrentPage != 0 {
		IndexPage.CurrentPage--
	}

	var (
		getBooks []*model.Book
		err      error
	)
	cart := &model.Cart{}

	if IndexPage.Max != 0 {
		getBooks, err = IndexPage.GetBooksByPrice()
	} else {
		getBooks, err = IndexPage.GetBooks()
	}

	IndexPage.GetPages()

	session_gettemp, _ := r.Cookie("user_session")
	session_temp := session_gettemp.String()

	IsLogin, Admin := false, false
	username_temp := ""

	if session_temp != "" {

		currentsession := model.CheckLogin(session_temp[13:])

		if currentsession.User_name != "" {
			IsLogin = true
			username_temp = currentsession.User_name
			cart = model.GetCartByUserId(currentsession.User_id)
		}
		Admin = model.CheckAdmin(currentsession.User_id)

	}
	if err != nil {
		fmt.Println("controller's getbooks error")
		return
	} else {
		t.Execute(w, struct {
			Books     []*model.Book
			Page      model.Pages
			IsLogin   bool
			User_name string
			Admin     bool
			Cart      *model.Cart
		}{
			Books:     getBooks,
			Page:      *IndexPage,
			IsLogin:   IsLogin,
			User_name: username_temp,
			Cart:      cart,
			Admin:     Admin,
		})
	}

}
func Login(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/pages/user/login.html"))
	t.Execute(w, struct {
		ClueStr string
	}{
		ClueStr: "请输入用户名和密码",
	})
}

func Regist(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/pages/user/regist.html"))
	t.Execute(w, nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session_gettemp, _ := r.Cookie("user_session")
	session_temp := session_gettemp.String()
	model.Logout(session_temp[13:])
	session_gettemp.MaxAge = -1

	http.SetCookie(w, session_gettemp)

	Index(w, r)
}

func main() {
	//solve static

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("src/static"))))
	//http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("src/pages"))))

	http.HandleFunc("/index", Index)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/logout", Logout)
	http.HandleFunc("/regist", Regist)

	http.HandleFunc("/login_verify", controller.Login_verify)
	http.HandleFunc("/regist_verify", controller.Regist_verify)

	http.HandleFunc("/getbooks", controller.GetBooks)
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	http.HandleFunc("/updateBook", controller.UpdateBook)
	http.HandleFunc("/toupdateBook", controller.ToupdateBook)

	http.HandleFunc("/cart", controller.Getcart)
	http.HandleFunc("/addBooktoCart", controller.AddBooktoCart)

	http.HandleFunc("/getorders", controller.GetOrders)
	http.HandleFunc("/get_order_detail", controller.GetOrderDetail)
	http.HandleFunc("/checkout", controller.Checkout)
	http.HandleFunc("/myorders", controller.MyOrders)
	http.HandleFunc("/clearcart", controller.Clearcart)
	http.HandleFunc("/sendoutgoods", controller.Sendoutgoods)
	http.HandleFunc("/confirmOrder", controller.ConfirmOrder)

	http.ListenAndServe(":8080", nil)
}
