package models

import "database/sql"

func MakeDbSalesPerson() *DbSalesPerson {
	return &DbSalesPerson{}
}

type (
	SalesPerson struct {
		Id                   int     `json:"id"`
		Name                 string  `json:"name"`
		SalesPersonOrderList []Order `json:"salesPersonOrderList"`
		EmployeeId           int     `json:"employeeId"`
	}

	DbSalesPerson struct {
		Id                   *sql.NullInt32  `json:"id"`
		Name                 *sql.NullString `json:"name"`
		SalesPersonOrderList []DbOrder       `json:"salesPersonOrderList"`
		EmployeeId           *sql.NullInt32  `json:"employeeId"`
	}
)

func GetSalesPersonsTobeDeletedAndAdded(existingSalesPerson, newSalesPerson []SalesPerson) (salesPersonsToBeUpdated, salesPersonsToBeAdded, salesPersonsToBeDeleted []SalesPerson) {

	m := make(map[int]bool)
	idSalesPersonMap := make(map[int]SalesPerson)

	for _, item := range existingSalesPerson {
		m[item.Id] = true
		idSalesPersonMap[item.Id] = item
	}
	for _, item := range newSalesPerson {
		if item.Id == 0 {
			salesPersonsToBeAdded = append(salesPersonsToBeAdded, item)
		} else if _, ok := m[item.Id]; ok {
			salesPersonsToBeUpdated = append(salesPersonsToBeUpdated, item)
			delete(m, item.Id)
		} else {
			salesPersonsToBeDeleted = append(salesPersonsToBeDeleted, item)
		}

	}

	for k, _ := range m {
		salesPersonsToBeDeleted = append(salesPersonsToBeDeleted, idSalesPersonMap[k])
	}
	return

}

func FromDbSalesPerson(dbSalesPerson *DbSalesPerson) (salesPerson *SalesPerson) {
	salesPerson = &SalesPerson{}
	if dbSalesPerson.Id != nil && dbSalesPerson.Id.Valid {
		salesPerson.Id = int(dbSalesPerson.Id.Int32)
	}
	if dbSalesPerson.Name != nil && dbSalesPerson.Name.Valid {
		salesPerson.Name = dbSalesPerson.Name.String
	}
	if dbSalesPerson.EmployeeId != nil && dbSalesPerson.EmployeeId.Valid {
		salesPerson.EmployeeId = int(dbSalesPerson.EmployeeId.Int32)
	}
	return
}

func ToDbSalesPerson(salesPerson *SalesPerson) (salesPersonDb *DbSalesPerson) {
	salesPersonDb = &DbSalesPerson{}
	salesPersonDb.Id = getNullInt(salesPerson.Id)
	salesPersonDb.Name = getNullString(salesPerson.Name)
	salesPersonDb.EmployeeId = getNullInt(salesPerson.EmployeeId)
	return
}
