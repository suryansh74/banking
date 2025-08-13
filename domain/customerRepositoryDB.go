package domain

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

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

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	// CONNECTION ESTABLISHMENT
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWORD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	logger.Info("value passed")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client}
}
