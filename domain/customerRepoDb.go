package domain

import (
	"Banking/err"
	"Banking/logger"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

type CustomerRepoDb struct {
	client *sqlx.DB
}

func (d CustomerRepoDb) FindAll(status string) ([]Customer, *err.AppError) {
	customers := make([]Customer, 0)
	var error error
	if status == "" {
		findAllSql := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers"
		error = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers where status = ?"
		error = d.client.Select(&customers, findAllSql, status)
	}

	if error != nil {
		logger.Error(error.Error())
		if error == sql.ErrNoRows {
			return nil, err.NotFoundError("customer not found")
		} else {
			return nil, err.UnexpectedError("unexpected database error")
		}
	}

	//convert (scan) rows into domainObj using sqlx
	//e := sqlx.StructScan(customerList, &customers)
	//if e != nil {
	//	logger.Error(e.Error())
	//	return nil, err.UnexpectedError("unexpected database error")
	//}

	//convert (scan) rows into domainObj manually
	//for customerList.Next() {
	//	var cust Customer
	//	error := customerList.Scan(&cust.Id, &cust.Name, &cust.DateOfBirth, &cust.City, &cust.Zipcode, &cust.Status)
	//	if error != nil {
	//		logger.Error(error.Error())
	//		return nil, err.UnexpectedError("unexpected database error")
	//	}
	//	customers = append(customers, cust)
	//}

	return customers, nil
}

func (d CustomerRepoDb) ById(customerId string) (*Customer, *err.AppError) {
	var customer Customer
	findByIdSql := "SELECT customer_id, name, date_of_birth, city, zipcode, status FROM customers WHERE customer_id = ?"
	error := d.client.Get(&customer, findByIdSql, customerId)

	if error != nil {
		logger.Error(error.Error())
		if error == sql.ErrNoRows {
			return nil, err.NotFoundError("customer not found with id " + customerId)
		} else {
			return nil, err.UnexpectedError("unexpected database error")
		}
	}

	return &customer, nil
}

func NewCustomerRepo() CustomerRepoDb {
	client, err := sqlx.Open("mysql", "rakshit:rakshit@(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepoDb{client}
}
