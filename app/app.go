package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/suryansh74/banking/domain"
	"github.com/suryansh74/banking/service"
)

func Start() {
	router := mux.NewRouter()
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Printf("Starting server at %s:%s\n", address, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
