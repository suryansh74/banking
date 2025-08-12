// Package domain: port for customer domain
package domain

import "github.com/suryansh74/banking/errs"

type Customer struct {
	ID          string
	Name        string
	City        string
	ZipCode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ByID(string) (*Customer, *errs.AppError)
}
