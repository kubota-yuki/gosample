package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

// func index(w http.ResponseWriter, r *http.Request) {
// 	t, _ := template.ParseFiles("template.html")
// 	t.Execute(w, nil)
// }

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":postgres://sfvermagpqsqkg:3b1bf4a863bf4e5c19a4282babe673aad32f97ad76e9e2f7459d461488bda821@ec2-34-225-159-178.compute-1.amazonaws.com:5432/db2tj1kv4jlep0", nil)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/form", form)
	http.HandleFunc("/result", result)
	http.HandleFunc("/make", make)
	http.HandleFunc("/", dbtest)
	http.HandleFunc("/init", del)
	http.HandleFunc("/create", create)
}

type User struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

func dbtest(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "user=sfvermagpqsqkg password=3b1bf4a863bf4e5c19a4282babe673aad32f97ad76e9e2f7459d461488bda821 dbname=db2tj1kv4jlep0 host=localhost port=5432 sslmode=disable")
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

	rows, err := db.Query("delete * from testtable")
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

func form(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("make.html")
	t.Execute(w, nil)
	db, err := sql.Open("postgres", "user=postgres password=yuki2170286 dbname=testdb host=localhost port=5432 sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("create table testtable (id varchar(10),name varchar(20),population int,date_mod date)")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

}
