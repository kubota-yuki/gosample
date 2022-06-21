package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

// func index(w http.ResponseWriter, r *http.Request) {
// 	t, _ := template.ParseFiles("template.html")
// 	t.Execute(w, nil)
// }

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template.html")
	t.Execute(w, nil)

}

func del(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("dbcreate.html")
	t.Execute(w, nil)

}

func result(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	hoge := r.FormValue("hoge")
	t, _ := template.ParseFiles("template2.html")
	t.Execute(w, hoge)

}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/form", form)
	http.HandleFunc("/result", result)
	http.HandleFunc("/make", make)
	http.HandleFunc("/", dbtest)
	http.HandleFunc("/init", del)
	http.HandleFunc("/create", create)
	server.ListenAndServe()
}

type User struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func dbtest(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "user=postgres password=yuki2170286 dbname=testdb host=localhost port=5432 sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT * FROM testtable") //
	if err != nil {
		panic(err.Error())
	}

	var u User
	for rows.Next() {
		rows.Scan(&u.ID, &u.Name)
		fmt.Fprintln(w, u.ID, u.Name)
	}
	defer rows.Close()
}

func make(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("make.html")

	db, err := sql.Open("postgres", "user=postgres password=yuki2170286 dbname=testdb host=localhost port=5432 sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("truncate table testtable")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()
	t.Execute(w, nil)

}

func create(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("postgres", "user=postgres password=yuki2170286 dbname=testdb host=localhost port=5432 sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("INSERT INTO testtable(id,name) VALUES(" + r.FormValue("id") + ",'" + r.FormValue("name") + "')")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()
	t, _ := template.ParseFiles("success.html")
	t.Execute(w, nil)

}
