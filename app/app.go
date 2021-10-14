package app

import (
	"log"
	"net/http"

	"github.com/andiahmads/bangking-service/domain"
	"github.com/andiahmads/bangking-service/service"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customer", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
	
}
