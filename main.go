package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func connectDb() *sql.DB {
	connection := "user=store dbname=one_store password=he20580020he host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := connectDb()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{Name: "T-shirt", Description: "Azure", Price: 49, Quantity: 5},
		{"Shoes", "Leather", 99, 3},
		{"Hat", "Beachwear", 23, 10},
	}
	temp.ExecuteTemplate(w, "Index", products)
}
