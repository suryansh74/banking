package domain

import (
	"database/sql"
	"errors"
	"log"

	"github.com/suryansh74/banking/errs"
	"github.com/suryansh74/banking/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {
	// SQL QUERY
	findAllSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.client.Query(findAllSQL)
	if err != nil {
		log.Println("Error while querying customer table" + err.Error())
		return nil, err
	}

	defer rows.Close()

	customers := make([]Customer, 0)
	err = sqlx.StructScan(rows, &customers)
	if err != nil {
		logger.Error("Error while scanning customers" + err.Error())
		return nil, err
	}
	return customers, nil
}

func (d CustomerRepositoryDB) ByID(id string) (*Customer, *errs.AppError) {
	customerSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	var customer Customer
	err := d.client.Get(&customer, customerSQL, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning customer:" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected error")
		}
	}
	return &customer, nil
}

func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{dbClient}
}
