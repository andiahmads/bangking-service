package domain

import (
	"database/sql"

	"github.com/andiahmads/bangking-service/helpers"
	"github.com/andiahmads/bangking-service/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *helpers.AppError) {
	// var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAllSql := "SELECT * FROM customers"
		err = d.client.Select(&customers, findAllSql)
		// rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := "SELECT * FROM customers"
		// rows, err = d.client.Query(findAllSql, status)
		err = d.client.Select(&customers, findAllSql, status)

	}

	if err != nil {
		logger.Error("Error while querying customer table: " + err.Error())
		return nil, helpers.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil

}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *helpers.AppError) {
	customerSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id = ?"

	// rows := d.client.QueryRow(customerSql, id)
	var c Customer
	err := d.client.Get(&c, customerSql, id)
	// err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helpers.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customers" + err.Error())
			return nil, helpers.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil

}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {

	return CustomerRepositoryDb{dbClient}
}
