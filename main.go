package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
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
