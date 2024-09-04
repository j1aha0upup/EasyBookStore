package model_test

// func TestAddUser(t *testing.T) {
// 	user := model.User{
// 		Username: "admin",
// 		Password: "admin",
// 		Email:    "admin@admin.com",
// 	}
// 	err := user.AddUser()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func TestCheckUsername(t *testing.T) {
// 	username := "admi1n"

// 	model.CheckUsername(username)
// }

// func TestLoginVerify(t *testing.T) {
// 	model.Login_verify("admin", "admi1n")
// }

// func TestGetBook(t *testing.T) {
// 	temp, _ := model.GetBooks()

// 	for i, v := range temp {
// 		fmt.Printf("No.%d book is %v\n", i, v)
// 	}
// }

// func TestAddBook(t *testing.T) {
// 	a := model.Book{
// 		Title:  "aa",
// 		Author: "dd",
// 		Price:  10,
// 		Sales:  100,
// 		Stock:  200,
// 	}
// 	model.AddBook(&a)

// }
// func TestDeleteBook(t *testing.T) {
// 	model.DeleteBook(32)

// }
// func TestGetBook(t *testing.T) {
// 	temp, _ := model.GetBook(22)
// 	fmt.Println(temp)
// }

// func TestGetSession(t *testing.T) {
// 	user_temp := &model.User{
// 		Username: "dd",
// 		Password: "dd",
// 		Id:       3,
// 		Email:    "d",
// 	}
// 	session_temp := model.Login(*user_temp)
// 	fmt.Println(session_temp)
// }

// func TestLogout(t *testing.T) {
// 	session_temp := model.Session{
// 		Uuid: "0e538417-c692-4ef2-8150-de3207cfd6b4",
// 	}
// 	model.Logout(session_temp)
// }
