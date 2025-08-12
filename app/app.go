package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/suryansh74/banking/domain"
	"github.com/suryansh74/banking/service"
)

func Start() {
	router := mux.NewRouter()
	// WIRING for domain and service
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
