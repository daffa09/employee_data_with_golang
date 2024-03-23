package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/file"
	_ "github.com/go-sql-driver/mysql"
)

type People struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int64  `json:"age"`
	Email string `json:"email"`
	Phone int64  `json:"phone"`
}

func Routes() {
	router := routing.New()
	router.Get("/", file.Content("views/index.html"))
	router.Get("/tambah", file.Content("views/tambah.html"))

	http.Handle("/", router)
	http.Handle("/tambah", router)
}

var db *sql.DB
var err error
var tpl *template.Template

type Person struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int64  `json:"age"`
	Email string `json:"email"`
	Phone int64  `json:"phone"`
}

func init() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/learning")
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	tpl = template.Must(template.ParseGlob("views/*"))
}

func main() {

	Routes()

	styles := http.FileServer(http.Dir("./asset"))
	http.Handle("/asset/", http.StripPrefix("/asset/", styles))

	fmt.Println("Server is running at http://localhost:8090")
	http.ListenAndServe(":8090", nil)
}

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

	users := make([]Person, 0)
	for rows.Next() {
		usr := Person{}
		rows.Scan(&usr.ID, &usr.Name, &usr.Age, &usr.Email, &usr.Phone)
		users = append(users, usr)
	}
	log.Println(users)
	tpl.ExecuteTemplate(w, "index.html", users)
}
