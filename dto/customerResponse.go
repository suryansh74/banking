// Package dto: returning response data type object mapped to customer from domain
package dto

type CustomerResponse struct {
	ID          string `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code"`
	DateofBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}
