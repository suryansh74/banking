// Package domain: port for customer domain
package domain

import (
	"github.com/suryansh74/banking/dto"
	"github.com/suryansh74/banking/errs"
)

type Customer struct {
	ID          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ByID(string) (*Customer, *errs.AppError)
}

func (c Customer) statusAsText() string {
	statusText := "active"
	if c.Status == "0" {
		statusText = "inactive"
	}
	return statusText
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}
