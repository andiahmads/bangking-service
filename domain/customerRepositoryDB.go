package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/andiahmads/bangking-service/errs"
	"github.com/andiahmads/bangking-service/logger"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllSql := "select * from customers"
	// findAllSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers"
	log.Print(findAllSql)

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		logger.Error("Error while querying customer table: " + err.Error())
		return nil, err
	}

	customer := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

		if err != nil {
			logger.Error("Error while scanning customer table: " + err.Error())
			return nil, err
		}
		customer = append(customer, c)

	}
	return customer, nil

}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id = ?"

	rows := d.client.QueryRow(customerSql, id)
	var c Customer
	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customers" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil

}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:endi@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	client.SetMaxOpenConns(10)
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetConnMaxIdleTime(10)
	return CustomerRepositoryDb{client}
}
