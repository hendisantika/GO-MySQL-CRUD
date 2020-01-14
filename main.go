package GO_MySQL_CRUD_Demo

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error
var tpl *template.Template

type user struct {
	ID        int64
	Username  string
	Firstname string
	Lastname  string
	Password  []byte
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/userForm", userForm)
	http.HandleFunc("/createUsers", createUsers)
	http.HandleFunc("/editUsers", editUsers)
	http.HandleFunc("/deleteUsers", deleteUsers)
	http.HandleFunc("/updateUsers", updateUsers)
	log.Println("Server is up on 8080 port")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
