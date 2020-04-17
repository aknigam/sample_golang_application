package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"sample_golang_application/common"
	"sample_golang_application/models"
)

type CustomerRepository struct {
}

func (repo *CustomerRepository) GetCustomer(customerId int) (customer *models.Customer, err error) {
	dbcustomer := models.MakeDbCustomer()
	txn, err := Db.Begin()
	if err != nil {
		return
	}
	defer txn.Commit()
	err = Db.QueryRow(" SELECT id, name"+
		" FROM customer"+
		" WHERE (id = ?)", customerId).
		Scan(&dbcustomer.Id, &dbcustomer.Name)
	if err != nil {
		txn.Rollback()
	}
	return models.FromDbCustomer(dbcustomer), err
}

func (repo *CustomerRepository) DeleteCustomer(customerId int) (err error) {
	_, err = Db.Exec(" DELETE FROM customer"+
		" WHERE (id =  ?)", customerId)
	return
}

func (repo *CustomerRepository) UpdateCustomer(customer *models.Customer) (err error) {
	dbcustomer := models.ToDbCustomer(customer)
	_, err = Db.Exec(" UPDATE customer"+
		" SET name = ?"+
		" WHERE (id = ?)", dbcustomer.Name, dbcustomer.Id)
	if err != nil {
		common.Error.Println("Customer could not be updated ")
		return
	}
	return
}

func (repo *CustomerRepository) CreateCustomer(customer *models.Customer) (id int, err error) {
	dbcustomer := models.ToDbCustomer(customer)
	statement := " INSERT INTO customer" +
		"  (name)" +
		" VALUES (?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(dbcustomer.Name)

	if err != nil {
		common.Error.Println("Could not create Customer ", err)
		return
	}
	generatedId, err := result.LastInsertId()
	if err != nil {
		common.Error.Println("Could not create Customer ", err)
		return
	}
	return int(generatedId), err

}
