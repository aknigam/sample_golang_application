package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"sample_golang_application/common"
	"sample_golang_application/models"
)

type ProductRepository struct {
}

func (repo *ProductRepository) GetProduct(productId int) (product *models.Product, err error) {
	dbproduct := models.MakeDbProduct()
	txn, err := Db.Begin()
	if err != nil {
		return
	}
	defer txn.Commit()
	err = Db.QueryRow(" SELECT id, name"+
		" FROM product"+
		" WHERE (id = ?)", productId).
		Scan(&dbproduct.Id, &dbproduct.Name)
	if err != nil {
		txn.Rollback()
	}
	return models.FromDbProduct(dbproduct), err
}

func (repo *ProductRepository) DeleteProduct(productId int) (err error) {
	_, err = Db.Exec(" DELETE FROM product"+
		" WHERE (id =  ?)", productId)
	return
}

func (repo *ProductRepository) UpdateProduct(product *models.Product) (err error) {
	dbproduct := models.ToDbProduct(product)
	_, err = Db.Exec(" UPDATE product"+
		" SET name = ?"+
		" WHERE (id = ?)", dbproduct.Name, dbproduct.Id)
	if err != nil {
		common.Error.Println("Product could not be updated ")
		return
	}
	return
}

func (repo *ProductRepository) CreateProduct(product *models.Product) (id int, err error) {
	dbproduct := models.ToDbProduct(product)
	statement := " INSERT INTO product" +
		"  (name)" +
		" VALUES (?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(dbproduct.Name)

	if err != nil {
		common.Error.Println("Could not create Product ", err)
		return
	}
	generatedId, err := result.LastInsertId()
	if err != nil {
		common.Error.Println("Could not create Product ", err)
		return
	}
	return int(generatedId), err

}

func (repo *ProductRepository) GetAllOrderItemProducts(orderItemId int) (products []models.Product) {
	rows, err := Db.Query(" SELECT pb.id, pb.name"+
		" FROM product pb"+
		" JOIN order_item_products oipa on oipa.product_id = pb.id"+
		" WHERE (oipa.order_item_id = ?)", orderItemId)
	if err != nil {
		common.Error.Println("Could not find products ", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		product := models.Product{}
		err := rows.Scan(&product.Id, &product.Name)
		if err != nil {
			common.Error.Println("Could not find products ", err)
			break
		}
		products = append(products, product)
	}
	err = rows.Err()
	if err != nil {
		common.Error.Fatal(err)
	}
	return
}
