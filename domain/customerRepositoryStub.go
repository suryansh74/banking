// Package domain: One of the Adapter for port
package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Ronak", "Kota", "204412", "10-09-2002", "1"},
		{"1002", "Kartik", "Noida", "204412", "28-06-2002", "1"},
	}
	return CustomerRepositoryStub{customers}
}
