package models

import "database/sql"

func MakeDbOrderItem() *DbOrderItem {
	return &DbOrderItem{}
}

type (
	OrderItem struct {
		Id       int       `json:"id"`
		Name     string    `json:"name"`
		OrderId  int       `json:"orderId"`
		Products []Product `json:"products"`
	}

	DbOrderItem struct {
		Id       *sql.NullInt32  `json:"id"`
		Name     *sql.NullString `json:"name"`
		OrderId  *sql.NullInt32  `json:"orderId"`
		Products []DbProduct     `json:"products"`
	}
)

func GetOrderItemsTobeDeletedAndAdded(existingOrderItem, newOrderItem []OrderItem) (orderItemsToBeUpdated, orderItemsToBeAdded, orderItemsToBeDeleted []OrderItem) {

	m := make(map[int]bool)
	idOrderItemMap := make(map[int]OrderItem)

	for _, item := range existingOrderItem {
		m[item.Id] = true
		idOrderItemMap[item.Id] = item
	}
	for _, item := range newOrderItem {
		if item.Id == 0 {
			orderItemsToBeAdded = append(orderItemsToBeAdded, item)
		} else if _, ok := m[item.Id]; ok {
			orderItemsToBeUpdated = append(orderItemsToBeUpdated, item)
			delete(m, item.Id)
		} else {
			orderItemsToBeDeleted = append(orderItemsToBeDeleted, item)
		}

	}

	for k, _ := range m {
		orderItemsToBeDeleted = append(orderItemsToBeDeleted, idOrderItemMap[k])
	}
	return

}

func FromDbOrderItem(dbOrderItem *DbOrderItem) (orderItem *OrderItem) {
	orderItem = &OrderItem{}
	if dbOrderItem.Id != nil && dbOrderItem.Id.Valid {
		orderItem.Id = int(dbOrderItem.Id.Int32)
	}
	if dbOrderItem.Name != nil && dbOrderItem.Name.Valid {
		orderItem.Name = dbOrderItem.Name.String
	}
	return
}

func ToDbOrderItem(orderItem *OrderItem) (orderItemDb *DbOrderItem) {
	orderItemDb = &DbOrderItem{}
	orderItemDb.Id = getNullInt(orderItem.Id)
	orderItemDb.Name = getNullString(orderItem.Name)
	return
}
