package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"sample_golang_application/common"
	"sample_golang_application/models"
)

type OrderRepository struct {
}

func (repo *OrderRepository) GetOrder(orderId int) (order *models.Order, err error) {
	dborder := models.MakeDbOrder()
	txn, err := Db.Begin()
	if err != nil {
		return
	}
	defer txn.Commit()
	err = Db.QueryRow(" SELECT id, name, customer, sales_person"+
		" FROM table_order"+
		" WHERE (id = ?)", orderId).
		Scan(&dborder.Id, &dborder.Name, &dborder.Customer, &dborder.SalesPerson)
	if err != nil {
		txn.Rollback()
	}
	return models.FromDbOrder(dborder), err
}

func (repo *OrderRepository) DeleteOrder(orderId int) (err error) {
	_, err = Db.Exec(" DELETE FROM table_order"+
		" WHERE (id =  ?)", orderId)
	return
}

func (repo *OrderRepository) UpdateOrder(order *models.Order) (err error) {
	dborder := models.ToDbOrder(order)
	_, err = Db.Exec(" UPDATE table_order"+
		" SET name = ?, customer = ?, sales_person = ?"+
		" WHERE (id = ?)", dborder.Name, dborder.Customer, dborder.SalesPerson, dborder.Id)
	if err != nil {
		common.Error.Println("Order could not be updated ")
		return
	}
	return
}

func (repo *OrderRepository) CreateOrder(order *models.Order) (id int, err error) {
	dborder := models.ToDbOrder(order)
	statement := " INSERT INTO table_order" +
		"  (name, customer, sales_person)" +
		" VALUES (?, ?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(dborder.Name, dborder.Customer, dborder.SalesPerson)

	if err != nil {
		common.Error.Println("Could not create Order ", err)
		return
	}
	generatedId, err := result.LastInsertId()
	if err != nil {
		common.Error.Println("Could not create Order ", err)
		return
	}
	return int(generatedId), err

}

func (repo *OrderRepository) LinkOrderOrderItems(orderId int, orderItemId int) (err error) {
	_, err = Db.Exec(" UPDATE order_item"+
		" SET order_id= ?"+
		" WHERE (id= ?)", orderId, orderItemId)
	if err != nil {
		common.Error.Println("Failed to link orderItem with order ")
		panic(err)
	}
	return
}
func (repo *OrderRepository) UnlinkOrderOrderItems(orderId int, orderItemId int) (err error) {
	_, err = Db.Exec(" UPDATE order_item"+
		" SET order_id= null"+
		" WHERE (id= ?)", orderItemId)
	if err != nil {
		common.Error.Println("Failed to unlink orderItems from order ")
		panic(err)
	}
	return
}
func (repo *OrderRepository) UnlinkAllOrderOrderItems(orderId int) (err error) {
	_, err = Db.Exec(" UPDATE order_item"+
		" SET order_id= null"+
		" WHERE (order_id= ?)", orderId)
	if err != nil {
		common.Error.Println("Could not delete orderItems ")
		panic(err)
	}
	return
}
func (repo *OrderRepository) DeleteAllOrderOrderItems(orderId int) (err error) {
	_, err = Db.Exec("delete from order_item where order_id = ?", orderId)
	if err != nil {
		common.Error.Println("orderItems could not be deleted ")
		panic(err)
	}
	return
}
func (repo *OrderRepository) GetAllSalesPersonSalesPersonOrderList(salesPersonId int) (orders []models.Order) {
	rows, err := Db.Query(" SELECT id, name, customer, sales_person"+
		" FROM table_order"+
		" WHERE (sales_person = ?)", salesPersonId)
	if err != nil {
		common.Error.Println("Could not find orders ", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		order := models.Order{}
		err := rows.Scan(&order.Id, &order.Name, &order.Customer, &order.SalesPerson)
		if err != nil {
			common.Error.Println("Could not find orders ", err)
			break
		}
		orders = append(orders, order)
	}
	err = rows.Err()
	if err != nil {
		common.Error.Fatal(err)
	}
	return
}
