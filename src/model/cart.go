package model

import (
	"book_store_demo/src/utils"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	Id            string
	InCartItems   []*InCartItem
	TotalQuantity int
	TotalAmount   float64
	User_id       int
}

type InCartItem struct {
	Id       int
	Books    Book
	Book_id  int
	Quantity int
	Amount   float64
	Cart_id  string
}

func GetCartByUserId(user_id int) *Cart {
	result := &Cart{}
	sqlStr := "select id,totalQuantity,totalAmount,user_id from carts where user_id=?"
	queryResult := utils.Db.QueryRow(sqlStr, user_id)
	queryResult.Scan(&result.Id, &result.TotalQuantity, &result.TotalAmount, &result.User_id)

	return result
}

func GetCartByCartId(cart_id string) *Cart {
	result := &Cart{}
	sqlStr := "select id,totalQuantity,totalAmount,user_id from carts where id=?"
	queryResult := utils.Db.QueryRow(sqlStr, cart_id)
	queryResult.Scan(&result.Id, &result.TotalQuantity, &result.TotalAmount, &result.User_id)

	return result
}

func (cart *Cart) GetInCartItem() {
	sqlStr := "select id,book_id,quantity,amount,cart_id from in_cart_item where cart_id=?"
	queryResult, err := utils.Db.Query(sqlStr, cart.Id)
	if err != nil {
		fmt.Println("func GetInCartItem err")
	}

	for queryResult.Next() {
		var cartItem_temp InCartItem

		queryResult.Scan(&cartItem_temp.Id, &cartItem_temp.Book_id, &cartItem_temp.Quantity, &cartItem_temp.Amount, &cartItem_temp.Cart_id)

		cartItem_temp.Books, err = GetBookById(cartItem_temp.Book_id)
		if err != nil {
			fmt.Println("func GetInCartItem error")
		}

		cart.InCartItems = append(cart.InCartItems, &cartItem_temp)
	}
}

func (cart *Cart) AddBooktoUserCartbyCartId(bookid int) {
	cart.GetInCartItem()
	flagDontHas := true

	for _, v := range cart.InCartItems {
		if v.Book_id == bookid {
			flagDontHas = false
			v.Quantity++
			UpdateInCartItemByid(v.Quantity, v.Id)
		}
	}
	if flagDontHas {
		sqlStr := "insert into in_cart_item(book_id,quantity,cart_id) values(?,?,?)"
		utils.Db.Exec(sqlStr, bookid, 1, cart.Id)
	}
}

func UpdateInCartItemByid(quantity int, id int) {
	sqlStr := "update in_cart_item set quantity=? where id=?"
	utils.Db.Exec(sqlStr, quantity, id)
}

func (cart *Cart) UpdateAmount() {
	cart.TotalAmount, cart.TotalQuantity = 0, 0
	for _, v := range cart.InCartItems {
		v.Amount = v.Books.Price * float64(v.Quantity)
		sqlStr := "update in_cart_item set Amount=? where id=?"
		utils.Db.Exec(sqlStr, v.Amount, v.Id)
		cart.TotalAmount += v.Amount
		cart.TotalQuantity += v.Quantity
	}
	sqlStr := "update carts set totalQuantity=?,totalAmount=? where id=?"
	utils.Db.Exec(sqlStr, cart.TotalQuantity, cart.TotalAmount, cart.Id)
}

func (cart *Cart) CreateCart() {
	cart.Id = uuid.New().String()
	cart.TotalQuantity = 0

	sqlStr := "insert into carts(id,totalQuantity,user_id) values(?,?,?)"
	utils.Db.Exec(sqlStr, cart.Id, cart.TotalQuantity, cart.User_id)
}

func (cart *Cart) AddOrder() string {
	tempOrderid := uuid.NewString()

	Order_date := time.Now()
	// Order_datestring = Order_date.String()

	for _, v := range cart.InCartItems {
		sqlstr := "INSERT INTO order_item(book_id,quantity,amount,order_id) VALUES(?,?,?,?)"
		_, err := utils.Db.Exec(sqlstr, v.Book_id, v.Quantity, v.Amount, tempOrderid)
		if err != nil {
			fmt.Println("AddOrder err")
		}

		UpdateBookStock(v.Book_id, v.Quantity)
	}
	sqlstr := "INSERT INTO `order`(id,user_id,order_flag,order_date,total_quantity,total_amount) VALUES(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlstr, tempOrderid, cart.User_id, 0, Order_date, cart.TotalQuantity, cart.TotalAmount)
	if err != nil {
		fmt.Println("AddOrder err")
	}
	return tempOrderid
}

func (cart *Cart) DeleteThisCart() {
	sqlStr := "DELETE FROM carts where id=?"
	_, err := utils.Db.Exec(sqlStr, cart.Id)

	if err != nil {
		fmt.Println("DeleteThisCart err")
	}
	sqlStr = "DELETE FROM in_cart_item where cart_id=?"
	_, err = utils.Db.Exec(sqlStr, cart.Id)

	if err != nil {
		fmt.Println("DeleteThisCart err")
	}
}
