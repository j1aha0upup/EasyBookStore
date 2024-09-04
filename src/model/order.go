package model

import (
	"book_store_demo/src/utils"
	"fmt"
	"time"
)

type Order struct {
	Id               string
	User_id          int
	Order_flag       int //订单状态 0为未发货 1为已发货 2为已完成
	Flag             string
	Order_date       time.Time
	Order_datestring string
	TotalQuantity    int
	TotalAmount      float64
	OrderItem        []*Order_item
}

type Order_item struct {
	Id       int
	Book     Book
	Quantity int
	Amount   float64
	Order_id string
}

func (page *Pages) GetAllOrders() (result []*Order, err error) {
	sqlStr := "select id,user_id,order_flag,total_quantity,total_amount,order_date from `order` limit ?,?"
	queryResult, err := utils.Db.Query(sqlStr, page.CurrentPage*page.EveryPageRecordCount, page.EveryPageRecordCount)

	if err != nil {
		fmt.Println("GetAllOrders Error")
		return nil, err
	}
	for queryResult.Next() {
		temp := Order{}
		var temporderdate string
		queryResult.Scan(&temp.Id, &temp.User_id, &temp.Order_flag, &temp.TotalQuantity, &temp.TotalAmount, &temporderdate)
		temp.Order_date, _ = time.Parse("2006-01-02 15:04:05", temporderdate)

		temp.OrderStatus()
		result = append(result, &temp)
	}
	sqlStr = "select COUNT(*) from `order`"
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&page.TotalRecords)
	return result, nil
}

func (order *Order) OrderStatus() {
	if order.Order_flag == 0 {
		order.Flag = "未发货"
	} else if order.Order_flag == 1 {
		order.Flag = "已发货"
	} else if order.Order_flag == 2 {
		order.Flag = "已完成"
	}
}

func (order *Order) GetOrderById() {
	sqlStr := "select id,user_id,order_flag,total_quantity,total_amount,order_date from `order` WHERE id=?"
	row := utils.Db.QueryRow(sqlStr, order.Id)
	var temporderdate string

	row.Scan(&order.Id, &order.User_id, &order.Order_flag, &order.TotalQuantity, &order.TotalAmount, &temporderdate)
	order.Order_date, _ = time.Parse("2006-01-02 15:04:05", temporderdate)
	order.OrderStatus()
}

func (order *Order) GetOrderDetail() {
	sqlStr := "select * from order_item WHERE order_id=?"
	queryResult, err := utils.Db.Query(sqlStr, order.Id)
	if err != nil {
		fmt.Println("GetOrderDetail error")
		fmt.Println(err.Error())
	}
	for queryResult.Next() {
		item := &Order_item{}

		queryResult.Scan(&item.Id, &item.Book.ID, &item.Quantity, &item.Amount, &item.Order_id)
		item.Book, _ = GetBookById(item.Book.ID)

		order.OrderItem = append(order.OrderItem, item)

	}

}

func (page *Pages) GetOrdersByUserId(user_id int) (result []*Order, err error) {
	sqlStr := "select id,user_id,order_flag,total_quantity,total_amount,order_date from `order` where user_id=? limit ?,?"
	queryResult, err := utils.Db.Query(sqlStr, user_id, page.CurrentPage*page.EveryPageRecordCount, page.EveryPageRecordCount)

	if err != nil {
		fmt.Println("GetAllOrders Error")
		return nil, err
	}
	for queryResult.Next() {
		temp := Order{}
		var temporderdate string

		queryResult.Scan(&temp.Id, &temp.User_id, &temp.Order_flag, &temp.TotalQuantity, &temp.TotalAmount, &temporderdate)
		temp.Order_date, _ = time.Parse("2006-01-02 15:04:05", temporderdate)
		temp.OrderStatus()
		result = append(result, &temp)
	}
	sqlStr = "select COUNT(*) from `order` where user_id=?"
	row := utils.Db.QueryRow(sqlStr, user_id)
	row.Scan(&page.TotalRecords)
	return result, nil
}

func (order *Order) SendoutgoodsById() {
	sqlStr := "UPDATE `order` SET order_flag=? where id=?"
	utils.Db.Exec(sqlStr, 1, order.Id)
}

func (order *Order) ConfirmOrderById() {
	sqlStr := "UPDATE `order` SET order_flag=? where id=?"
	utils.Db.Exec(sqlStr, 2, order.Id)
}
