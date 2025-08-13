package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/suryansh74/banking/errs"
	"github.com/suryansh74/banking/logger"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"
	result, err := d.client.Exec(sqlInsert, a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error Occuring while creating customer" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Error from Database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error Occuring while extracting id for new account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Error from Database")
	}
	a.AccountID = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
