package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name, City, Zipcode string
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")

}

func getAllCustomer(w http.ResponseWriter, r *http.Request) {

	customer := []Customer{
		{"andi", "jakarta", "203"},
		{"joko", "jakarta selatan", "203"},
	}
	json.NewEncoder(w).Encode(customer)
}
