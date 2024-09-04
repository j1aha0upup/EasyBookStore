package controller

import (
	"book_store_demo/src/model"
	"html/template"
	"net/http"
	"strconv"
)

func Getcart(w http.ResponseWriter, r *http.Request) {
	session_gettemp, _ := r.Cookie("user_session")
	session_temp := session_gettemp.String()
	IsLogin, CanCheck := false, false
	username_temp := ""
	cart := &model.Cart{}
	if session_temp != "" {
		currentsession := model.CheckLogin(session_temp[13:])
		if currentsession.User_name != "" {
			IsLogin = true
			username_temp = currentsession.User_name
			cart = model.GetCartByUserId(currentsession.User_id)
			cart.GetInCartItem()

			cart.UpdateAmount()
		}
	}

	if len(cart.InCartItems) != 0 {
		CanCheck = true
	}

	t := template.Must(template.ParseFiles("src/pages/cart/cart.html"))
	t.Execute(w, struct {
		IsLogin       bool
		Username_temp string
		Cart          *model.Cart
		CanCheck      bool
	}{
		IsLogin:       IsLogin,
		Username_temp: username_temp,
		Cart:          cart,
		CanCheck:      CanCheck,
	})
}

func AddBooktoCart(w http.ResponseWriter, r *http.Request) {
	session_gettemp, _ := r.Cookie("user_session")
	session_temp := session_gettemp.String()

	if session_temp != "" {
		currentsession := model.CheckLogin(session_temp[13:])
		if currentsession.User_name != "" {

			cart := model.GetCartByUserId(currentsession.User_id)

			if cart.Id == "" {
				cart.User_id = currentsession.User_id
				cart.CreateCart()
			}

			bookid, _ := strconv.Atoi(r.PostFormValue("bookId"))
			cart.AddBooktoUserCartbyCartId(bookid)
			book, _ := model.GetBookById(bookid)
			w.Write([]byte("add " + book.Title + " success~!"))

		} else {
			w.Write([]byte("Please login~!"))
		}
	} else {
		w.Write([]byte("Please login~!"))
	}
}

func Clearcart(w http.ResponseWriter, r *http.Request) {
	cartId := r.FormValue("cartId")
	cart := &model.Cart{
		Id: cartId,
	}
	cart.DeleteThisCart()
	Getcart(w, r)
}
