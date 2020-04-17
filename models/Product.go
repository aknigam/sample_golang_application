package models

import "database/sql"

func MakeDbProduct() *DbProduct {
	return &DbProduct{}
}

type (
	Product struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	DbProduct struct {
		Id   *sql.NullInt32  `json:"id"`
		Name *sql.NullString `json:"name"`
	}
)

func GetProductsTobeDeletedAndAdded(existingProduct, newProduct []Product) (productsToBeUpdated, productsToBeAdded, productsToBeDeleted []Product) {

	m := make(map[int]bool)
	idProductMap := make(map[int]Product)

	for _, item := range existingProduct {
		m[item.Id] = true
		idProductMap[item.Id] = item
	}
	for _, item := range newProduct {
		if item.Id == 0 {
			productsToBeAdded = append(productsToBeAdded, item)
		} else if _, ok := m[item.Id]; ok {
			productsToBeUpdated = append(productsToBeUpdated, item)
			delete(m, item.Id)
		} else {
			productsToBeDeleted = append(productsToBeDeleted, item)
		}

	}

	for k, _ := range m {
		productsToBeDeleted = append(productsToBeDeleted, idProductMap[k])
	}
	return

}

func FromDbProduct(dbProduct *DbProduct) (product *Product) {
	product = &Product{}
	if dbProduct.Id != nil && dbProduct.Id.Valid {
		product.Id = int(dbProduct.Id.Int32)
	}
	if dbProduct.Name != nil && dbProduct.Name.Valid {
		product.Name = dbProduct.Name.String
	}
	return
}

func ToDbProduct(product *Product) (productDb *DbProduct) {
	productDb = &DbProduct{}
	productDb.Id = getNullInt(product.Id)
	productDb.Name = getNullString(product.Name)
	return
}
