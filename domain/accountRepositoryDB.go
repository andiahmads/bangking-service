package domain

import (
	"strconv"

	"github.com/andiahmads/bangking-service/helpers"
	"github.com/andiahmads/bangking-service/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *helpers.AppError) {
	sqlInser := "INSERT INTO accounts (customer_id,opening_date,account_type,amount,status) values (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInser, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating Account" + err.Error())
		return nil, helpers.NewUnexpectedError("Unexpected Error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new Account" + err.Error())
		return nil, helpers.NewUnexpectedError("Unexpected Error")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}

}
