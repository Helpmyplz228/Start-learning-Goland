package main

import (
	"fmt"
	"net/http"
)

// Создание базавой структуры
type User struct {
	name                  string
	age                   uint16
	money                 int16
	avg_grades, happiness float64
}

// %s - строка, %d - число
// функция для вывода инфы из структуры
func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has money"+
		"equal: %d", u.name, u.age, u.money)
}

// * - ссылка на объект а не создание его кпии
// Функция для требования нового мяни
func (u *User) setNewName(newName string) {
	u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	Bob := User{"Bob", 25, -50, 4.2, 0.8}
	Bob.setNewName("Alex") // Также помеял имя
	//Bob.name = "Alex" // Поменял имя
	fmt.Fprintf(w, Bob.getAllInfo())
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contacts page")
}

func HandleRec() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	//var Bob User = ...
	//Bob := User{name: "Bob", age: 25, money: -50, avg_grades: 4.2, happiness: 0.8}
	// дно и тоже!
	//Bob := User{"Bob", 25, -50, 4.2, 0.8}
	HandleRec()
}
