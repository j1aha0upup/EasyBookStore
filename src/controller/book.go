package controller

import (
	"book_store_demo/src/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	Pages := &model.Pages{
		EveryPageRecordCount: 6,
	}
	Pages.CurrentPage, _ = strconv.Atoi(r.FormValue("currentPage"))

	if Pages.CurrentPage != 0 {
		Pages.CurrentPage--
	}

	getBooks, err := Pages.GetBooks()
	Pages.GetPages()
	if err != nil {
		fmt.Println("controller's getbooks error")
		return
	} else {

		template.Must(template.ParseFiles("src/pages/manager/book_manager.html")).Execute(w, struct {
			Books []*model.Book
			Page  model.Pages
		}{
			Books: getBooks,
			Page:  *Pages,
		})
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	model.DeleteBook(id)
	GetBooks(w, r)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	Book, err := model.GetBookById(id)
	if err != nil {
		fmt.Println("controller updateBook func has err")
		return
	}

	template.Must(template.ParseFiles("src/pages/manager/book_update.html")).Execute(w, struct {
		Book model.Book
	}{
		Book: Book,
	})
}

func ToupdateBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PostFormValue("id"))
	title := r.PostFormValue("title")
	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
	author := r.PostFormValue("author")
	sales, _ := strconv.Atoi(r.PostFormValue("sales"))
	stock, _ := strconv.Atoi(r.PostFormValue("stock"))

	temp := model.Book{
		ID:     id,
		Title:  title,
		Price:  price,
		Author: author,
		Sales:  sales,
		Stock:  stock,
	}
	if temp.ID != 0 {
		model.UpdateBook(temp)
	} else {
		temp.AddBook()
	}

	GetBooks(w, r)
}
