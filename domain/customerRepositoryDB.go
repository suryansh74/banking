package domain

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/suryansh74/banking/errs"
	"github.com/suryansh74/banking/logger"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
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

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.ZipCode, &c.DateofBirth, &c.Status)
		if err != nil {
			logger.Error("Error while scanning customers" + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDB) ByID(id string) (*Customer, *errs.AppError) {
	customerSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	row := d.client.QueryRow(customerSQL, id)
	var c Customer
	err := row.Scan(&c.ID, &c.Name, &c.City, &c.ZipCode, &c.DateofBirth, &c.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning customer:" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	// CONNECTION ESTABLISHMENT
	client, err := sql.Open("mysql", "root:@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client}
}
