package models

import "database/sql"

func MakeDbCustomer() *DbCustomer {
	return &DbCustomer{}
}

type (
	Customer struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	DbCustomer struct {
		Id   *sql.NullInt32  `json:"id"`
		Name *sql.NullString `json:"name"`
	}
)

func GetCustomersTobeDeletedAndAdded(existingCustomer, newCustomer []Customer) (customersToBeUpdated, customersToBeAdded, customersToBeDeleted []Customer) {

	m := make(map[int]bool)
	idCustomerMap := make(map[int]Customer)

	for _, item := range existingCustomer {
		m[item.Id] = true
		idCustomerMap[item.Id] = item
	}
	for _, item := range newCustomer {
		if item.Id == 0 {
			customersToBeAdded = append(customersToBeAdded, item)
		} else if _, ok := m[item.Id]; ok {
			customersToBeUpdated = append(customersToBeUpdated, item)
			delete(m, item.Id)
		} else {
			customersToBeDeleted = append(customersToBeDeleted, item)
		}

	}

	for k, _ := range m {
		customersToBeDeleted = append(customersToBeDeleted, idCustomerMap[k])
	}
	return

}

func FromDbCustomer(dbCustomer *DbCustomer) (customer *Customer) {
	customer = &Customer{}
	if dbCustomer.Id != nil && dbCustomer.Id.Valid {
		customer.Id = int(dbCustomer.Id.Int32)
	}
	if dbCustomer.Name != nil && dbCustomer.Name.Valid {
		customer.Name = dbCustomer.Name.String
	}
	return
}

func ToDbCustomer(customer *Customer) (customerDb *DbCustomer) {
	customerDb = &DbCustomer{}
	customerDb.Id = getNullInt(customer.Id)
	customerDb.Name = getNullString(customer.Name)
	return
}
