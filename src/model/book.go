package model

import (
	"book_store_demo/src/utils"
	"fmt"
)

type Book struct {
	ID        int
	Title     string
	Author    string
	Price     float64
	Sales     int
	Stock     int
	ImagePath string
}

func (page *Pages) GetBooks() (result []*Book, e error) {

	sqlStr := "select id,title,author,price,sales,stock,img_path from books limit ?,?"
	queryResult, err := utils.Db.Query(sqlStr, page.CurrentPage*page.EveryPageRecordCount, page.EveryPageRecordCount)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for queryResult.Next() {
		var book_temp Book

		queryResult.Scan(&book_temp.ID, &book_temp.Title, &book_temp.Author, &book_temp.Price, &book_temp.Sales, &book_temp.Stock, &book_temp.ImagePath)
		result = append(result, &book_temp)
	}

	sqlStr = "SELECT COUNT(*) FROM books"
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&page.TotalRecords)

	return result, nil
}

func (b *Book) AddBook() error {
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	b.ImagePath = "static/img/default.jpg"
	utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImagePath)
	return nil
}

func DeleteBook(id int) error {
	sqlStr := "delete from books where id=?"
	utils.Db.Exec(sqlStr, id)
	return nil
}

func GetBookById(id int) (result Book, err error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id=?"
	queryResult := utils.Db.QueryRow(sqlStr, id)
	queryResult.Scan(&result.ID, &result.Title, &result.Author, &result.Price, &result.Sales, &result.Stock, &result.ImagePath)
	return
}

func UpdateBook(temp Book) error {
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	utils.Db.Exec(sqlStr, temp.Title, temp.Author, temp.Price, temp.Sales, temp.Stock, temp.ID)
	return nil
}

func (pages *Pages) GetBooksByPrice() (result []*Book, e error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where (price BETWEEN ? AND ? ) limit ?,?"
	queryResult, err := utils.Db.Query(sqlStr, pages.Min, pages.Max, pages.CurrentPage*pages.EveryPageRecordCount, pages.EveryPageRecordCount)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for queryResult.Next() {
		var book_temp Book

		queryResult.Scan(&book_temp.ID, &book_temp.Title, &book_temp.Author, &book_temp.Price, &book_temp.Sales, &book_temp.Stock, &book_temp.ImagePath)
		result = append(result, &book_temp)
	}
	sqlStr = "SELECT COUNT(*) FROM books where (price BETWEEN ? AND ? )"
	row := utils.Db.QueryRow(sqlStr, pages.Min, pages.Max)
	row.Scan(&pages.TotalRecords)

	return result, nil
}

func UpdateBookStock(bookId int, quantity int) {
	book, err := GetBookById(bookId)
	if err != nil {
		fmt.Println("UpdateBookStock err")
	}
	stock := book.Stock - quantity

	sqlStr := "UPDATE books SET stock=? where id=?"
	utils.Db.Exec(sqlStr, stock, book.ID)
}
