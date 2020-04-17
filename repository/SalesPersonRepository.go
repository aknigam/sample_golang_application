package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"sample_golang_application/common"
	"sample_golang_application/models"
)

type SalesPersonRepository struct {
}

func (repo *SalesPersonRepository) GetSalesPerson(salesPersonId int) (salesPerson *models.SalesPerson, err error) {
	dbsalesPerson := models.MakeDbSalesPerson()
	txn, err := Db.Begin()
	if err != nil {
		return
	}
	defer txn.Commit()
	err = Db.QueryRow(" SELECT id, name, employee_id"+
		" FROM sales_person"+
		" WHERE (id = ?)", salesPersonId).
		Scan(&dbsalesPerson.Id, &dbsalesPerson.Name, &dbsalesPerson.EmployeeId)
	if err != nil {
		txn.Rollback()
	}
	return models.FromDbSalesPerson(dbsalesPerson), err
}

func (repo *SalesPersonRepository) DeleteSalesPerson(salesPersonId int) (err error) {
	_, err = Db.Exec(" DELETE FROM sales_person"+
		" WHERE (id =  ?)", salesPersonId)
	return
}

func (repo *SalesPersonRepository) UpdateSalesPerson(salesPerson *models.SalesPerson) (err error) {
	dbsalesPerson := models.ToDbSalesPerson(salesPerson)
	_, err = Db.Exec(" UPDATE sales_person"+
		" SET name = ?, employee_id = ?"+
		" WHERE (id = ?)", dbsalesPerson.Name, dbsalesPerson.EmployeeId, dbsalesPerson.Id)
	if err != nil {
		common.Error.Println("SalesPerson could not be updated ")
		return
	}
	return
}

func (repo *SalesPersonRepository) CreateSalesPerson(salesPerson *models.SalesPerson) (id int, err error) {
	dbsalesPerson := models.ToDbSalesPerson(salesPerson)
	statement := " INSERT INTO sales_person" +
		"  (name, employee_id)" +
		" VALUES (?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(dbsalesPerson.Name, dbsalesPerson.EmployeeId)

	if err != nil {
		common.Error.Println("Could not create SalesPerson ", err)
		return
	}
	generatedId, err := result.LastInsertId()
	if err != nil {
		common.Error.Println("Could not create SalesPerson ", err)
		return
	}
	return int(generatedId), err

}

func (repo *SalesPersonRepository) LinkSalesPersonSalesPersonOrderList(salesPersonId int, salesPersonOrderListId int) (err error) {
	_, err = Db.Exec(" UPDATE table_order"+
		" SET sales_person= ?"+
		" WHERE (id= ?)", salesPersonId, salesPersonOrderListId)
	if err != nil {
		common.Error.Println("Failed to link salesPersonOrderList with salesPerson ")
		panic(err)
	}
	return
}
func (repo *SalesPersonRepository) UnlinkSalesPersonSalesPersonOrderList(salesPersonId int, salesPersonOrderListId int) (err error) {
	_, err = Db.Exec(" UPDATE table_order"+
		" SET sales_person= null"+
		" WHERE (id= ?)", salesPersonOrderListId)
	if err != nil {
		common.Error.Println("Failed to unlink salesPersonOrderList from salesPerson ")
		panic(err)
	}
	return
}
func (repo *SalesPersonRepository) UnlinkAllSalesPersonSalesPersonOrderList(salesPersonId int) (err error) {
	_, err = Db.Exec(" UPDATE table_order"+
		" SET sales_person= null"+
		" WHERE (sales_person= ?)", salesPersonId)
	if err != nil {
		common.Error.Println("Could not delete salesPersonOrderList ")
		panic(err)
	}
	return
}
func (repo *SalesPersonRepository) DeleteAllSalesPersonSalesPersonOrderList(salesPersonId int) (err error) {
	_, err = Db.Exec("delete from table_order where sales_person = ?", salesPersonId)
	if err != nil {
		common.Error.Println("salesPersonOrderList could not be deleted ")
		panic(err)
	}
	return
}
