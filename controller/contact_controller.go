package controller

import (
	"Magang/data_karyawan/model"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
)

var tpl *template.Template
var db *sql.DB

//! mendapatkan semua data users
func GetAllPerson(w http.ResponseWriter, r *http.Request) {

	rows, e := db.Query(
		`SELECT *
		FROM contacts`)

	if e != nil {
		log.Println(e)
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	users := make([]model.Person, 0)
	for rows.Next() {
		usr := model.Person{}
		rows.Scan(&usr.ID, &usr.Name, &usr.Age, &usr.Email, &usr.Phone)
		users = append(users, usr)
	}
	log.Println(users)
	tpl.ExecuteTemplate(w, "index.html", users)
}

//! Membuat data user
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	db, _ := dbx.Open("mysql", "root:@tcp(127.0.0.1:3306)/learning")

	name := r.FormValue("nama")
	age, _ := strconv.ParseInt(r.FormValue("age")[0:], 10, 32)
	email := r.FormValue("email")
	phone, _ := strconv.ParseInt(r.FormValue("age")[0:], 10, 32)

	user := model.Person{
		Name:  name,
		Age:   age,
		Email: email,
		Phone: phone,
	}

	// INSERT INTO customer (name, email, status) VALUES ('example', 'test@example.com', 0)
	err := db.Model(&user).Insert()
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	fmt.Println("Berhasil input")
}
