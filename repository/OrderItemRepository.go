package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"sample_golang_application/common"
	"sample_golang_application/models"
)

type OrderItemRepository struct {
}

func (repo *OrderItemRepository) GetOrderItem(orderItemId int) (orderItem *models.OrderItem, err error) {
	dborderItem := models.MakeDbOrderItem()
	txn, err := Db.Begin()
	if err != nil {
		return
	}
	defer txn.Commit()
	err = Db.QueryRow(" SELECT id, name, order_id"+
		" FROM order_item"+
		" WHERE (id = ?)", orderItemId).
		Scan(&dborderItem.Id, &dborderItem.Name, &dborderItem.OrderId)
	if err != nil {
		txn.Rollback()
	}
	return models.FromDbOrderItem(dborderItem), err
}

func (repo *OrderItemRepository) DeleteOrderItem(orderItemId int) (err error) {
	_, err = Db.Exec(" DELETE FROM order_item"+
		" WHERE (id =  ?)", orderItemId)
	return
}

func (repo *OrderItemRepository) UpdateOrderItem(orderItem *models.OrderItem) (err error) {
	dborderItem := models.ToDbOrderItem(orderItem)
	_, err = Db.Exec(" UPDATE order_item"+
		" SET name = ?, order_id = ?"+
		" WHERE (id = ?)", dborderItem.Name, dborderItem.OrderId, dborderItem.Id)
	if err != nil {
		common.Error.Println("OrderItem could not be updated ")
		return
	}
	return
}

func (repo *OrderItemRepository) CreateOrderItem(orderItem *models.OrderItem) (id int, err error) {
	dborderItem := models.ToDbOrderItem(orderItem)
	statement := " INSERT INTO order_item" +
		"  (name, order_id)" +
		" VALUES (?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(dborderItem.Name, dborderItem.OrderId)

	if err != nil {
		common.Error.Println("Could not create OrderItem ", err)
		return
	}
	generatedId, err := result.LastInsertId()
	if err != nil {
		common.Error.Println("Could not create OrderItem ", err)
		return
	}
	return int(generatedId), err

}

func (repo *OrderItemRepository) LinkOrderItemProducts(orderItemId int, productId int) (err error) {
	_, err = Db.Exec("Insert into order_item_products ( order_item_id, product_id) values (  ? ,  ?)", orderItemId, productId)
	if err != nil {
		common.Error.Println("Failed to link product with orderItem ")
		panic(err)
	}
	return
}
func (repo *OrderItemRepository) UnlinkOrderItemProducts(orderItemId int, productId int) (err error) {
	_, err = Db.Exec("delete from order_item_products where order_item_id = ?  and product_id = ?", orderItemId, productId)
	if err != nil {
		common.Error.Println("Failed to unlink product with orderItem ")
		panic(err)
	}
	return
}
func (repo *OrderItemRepository) UnlinkAllOrderItemProducts(orderItemId int) (err error) {
	_, err = Db.Exec("delete from order_item_products where order_item_id = ? ", orderItemId)
	if err != nil {
		common.Error.Println("Could not delete products ")
		panic(err)
	}
	return
}
func (repo *OrderItemRepository) DeleteAllOrderItemProducts(orderItemId int) (err error) {
	_, err = Db.Exec("delete from product where id in (select product_id from order_item_products where order_item_id = ?)", orderItemId)
	if err != nil {
		common.Error.Println("products could not be deleted ")
		panic(err)
	}
	return
}
func (repo *OrderItemRepository) GetAllOrderOrderItems(orderId int) (orderItems []models.OrderItem) {
	rows, err := Db.Query(" SELECT id, name, order_id"+
		" FROM order_item"+
		" WHERE (order_id = ?)", orderId)
	if err != nil {
		common.Error.Println("Could not find orderItems ", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		orderItem := models.OrderItem{}
		err := rows.Scan(&orderItem.Id, &orderItem.Name, &orderItem.OrderId)
		if err != nil {
			common.Error.Println("Could not find orderItems ", err)
			break
		}
		orderItems = append(orderItems, orderItem)
	}
	err = rows.Err()
	if err != nil {
		common.Error.Fatal(err)
	}
	return
}
