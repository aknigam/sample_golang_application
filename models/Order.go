package models

import "database/sql"

func MakeDbOrder() *DbOrder {
	return &DbOrder{}
}

type (
	Order struct {
		Id          int         `json:"id"`
		Name        string      `json:"name"`
		Customer    int         `json:"customer"`
		OrderItems  []OrderItem `json:"orderItems"`
		SalesPerson int         `json:"salesPerson"`
	}

	DbOrder struct {
		Id          *sql.NullInt32  `json:"id"`
		Name        *sql.NullString `json:"name"`
		Customer    *sql.NullInt32  `json:"customer"`
		OrderItems  []DbOrderItem   `json:"orderItems"`
		SalesPerson *sql.NullInt32  `json:"salesPerson"`
	}
)

func GetOrdersTobeDeletedAndAdded(existingOrder, newOrder []Order) (ordersToBeUpdated, ordersToBeAdded, ordersToBeDeleted []Order) {

	m := make(map[int]bool)
	idOrderMap := make(map[int]Order)

	for _, item := range existingOrder {
		m[item.Id] = true
		idOrderMap[item.Id] = item
	}
	for _, item := range newOrder {
		if item.Id == 0 {
			ordersToBeAdded = append(ordersToBeAdded, item)
		} else if _, ok := m[item.Id]; ok {
			ordersToBeUpdated = append(ordersToBeUpdated, item)
			delete(m, item.Id)
		} else {
			ordersToBeDeleted = append(ordersToBeDeleted, item)
		}

	}

	for k, _ := range m {
		ordersToBeDeleted = append(ordersToBeDeleted, idOrderMap[k])
	}
	return

}

func FromDbOrder(dbOrder *DbOrder) (order *Order) {
	order = &Order{}
	if dbOrder.Id != nil && dbOrder.Id.Valid {
		order.Id = int(dbOrder.Id.Int32)
	}
	if dbOrder.Name != nil && dbOrder.Name.Valid {
		order.Name = dbOrder.Name.String
	}
	return
}

func ToDbOrder(order *Order) (orderDb *DbOrder) {
	orderDb = &DbOrder{}
	orderDb.Id = getNullInt(order.Id)
	orderDb.Name = getNullString(order.Name)
	return
}
