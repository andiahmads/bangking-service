package app

import (
	"encoding/json"
	"net/http"

	"github.com/andiahmads/bangking-service/service"
)

type Customer struct {
	Name, City, Zipcode string
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomer(w http.ResponseWriter, r *http.Request) {

	customers, _ := ch.service.GetAllCustomer()

	json.NewEncoder(w).Encode(customers)
}
