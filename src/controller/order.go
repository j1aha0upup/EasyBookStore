package controller

import (
	"book_store_demo/src/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	Pages := &model.Pages{
		EveryPageRecordCount: 6,
	}
	Pages.CurrentPage, _ = strconv.Atoi(r.FormValue("currentPage"))

	if Pages.CurrentPage != 0 {
		Pages.CurrentPage--
	}

	Orders, err := Pages.GetAllOrders()
	Pages.GetPages()
	if err != nil {
		fmt.Println("controller's GetOrders error")
		return
	} else {
		template.Must(template.ParseFiles("src/pages/manager/order_manager.html")).Execute(w, struct {
			Orders []*model.Order
			Page   model.Pages
		}{
			Orders: Orders,
			Page:   *Pages,
		})
	}
}

func GetOrderDetail(w http.ResponseWriter, r *http.Request) {
	order_id := r.FormValue("order_id")
	order := &model.Order{}
	order.Id = order_id
	order.GetOrderById()
	order.GetOrderDetail()

	template.Must(template.ParseFiles("src/pages/manager/order_detail.html")).Execute(w, struct {
		Order *model.Order
	}{
		Order: order,
	})
}

func Checkout(w http.ResponseWriter, r *http.Request) {
	cartId := r.FormValue("cartId")
	cart := model.GetCartByCartId(cartId)
	cart.GetInCartItem()
	orderId := cart.AddOrder()
	cart.DeleteThisCart()

	session_gettemp, _ := r.Cookie("user_session")
	session_temp := session_gettemp.String()
	IsLogin := false
	username_temp := ""
	if session_temp != "" {
		session_gettemp, _ := r.Cookie("user_session")
		session_temp := session_gettemp.String()
		currentsession := model.CheckLogin(session_temp[13:])
		if currentsession.User_name != "" {
			IsLogin = true
			username_temp = currentsession.User_name
		}
	}

	template.Must(template.ParseFiles("src/pages/cart/checkout.html")).Execute(w, struct {
		OrderId  string
		IsLogin  bool
		Username string
	}{
		OrderId:  orderId,
		IsLogin:  IsLogin,
		Username: username_temp,
	})
}

func MyOrders(w http.ResponseWriter, r *http.Request) {
	session_gettemp, _ := r.Cookie("user_session")
	session_temp := session_gettemp.String()
	IsLogin := false
	username_temp := ""
	currentsession := model.Session{}
	if session_temp != "" {
		session_gettemp, _ := r.Cookie("user_session")
		session_temp := session_gettemp.String()
		currentsession = model.CheckLogin(session_temp[13:])
		if currentsession.User_name != "" {
			IsLogin = true
			username_temp = currentsession.User_name
		}
	}
	Pages := &model.Pages{
		EveryPageRecordCount: 6,
	}
	Pages.CurrentPage, _ = strconv.Atoi(r.FormValue("currentPage"))

	if Pages.CurrentPage != 0 {
		Pages.CurrentPage--
	}

	Orders, err := Pages.GetOrdersByUserId(currentsession.User_id)
	if err != nil {
		fmt.Println("Myorders err")
	}
	Pages.GetPages()

	template.Must(template.ParseFiles("src/pages/manager/myorders.html")).Execute(w, struct {
		Orders   []*model.Order
		Page     model.Pages
		IsLogin  bool
		Username string
	}{
		IsLogin:  IsLogin,
		Username: username_temp,
		Orders:   Orders,
		Page:     *Pages,
	})
}

func Sendoutgoods(w http.ResponseWriter, r *http.Request) {
	order_id := r.FormValue("order_id")
	order := &model.Order{}
	order.Id = order_id
	order.GetOrderById()
	order.SendoutgoodsById()
	GetOrders(w, r)
}

func ConfirmOrder(w http.ResponseWriter, r *http.Request) {
	order_id := r.FormValue("order_id")
	order := &model.Order{}
	order.Id = order_id
	order.GetOrderById()
	order.ConfirmOrderById()
	MyOrders(w, r)
}
