package service

import (
	"sample_golang_application/errors"
	"sample_golang_application/models"
	"sample_golang_application/repository"
)

type SalesPersonService struct {
	SalesPersonRepo repository.SalesPersonRepository
	ProductRepo     repository.ProductRepository
	OrderItemRepo   repository.OrderItemRepository
	OrderRepo       repository.OrderRepository
}

func (service *SalesPersonService) GetSalesPerson(salesPersonId int) (salesPerson *models.SalesPerson, appErr *errors.AppError) {
	salesPerson, err := service.SalesPersonRepo.GetSalesPerson(salesPersonId)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "salesPerson not found", Code: -1}
		return
	}
	salesPersonOrderList := service.GetAllSalesPersonSalesPersonOrderList(salesPersonId)
	salesPerson.SalesPersonOrderList = salesPersonOrderList
	return
}

// ------------------ start get -----------------
func (service *SalesPersonService) GetSalesPersonSalesPersonOrderList(salesPersonId, orderId int) (salesPersonOrderList *models.Order, appErr *errors.AppError) {
	salesPersonOrderList, err := service.OrderRepo.GetOrder(orderId)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "salesPerson not found", Code: -1}
		return
	}
	if salesPersonOrderList.SalesPerson != salesPersonId {
		return nil, nil
	}
	orderItems := service.GetAllOrderOrderItems(orderId)
	salesPersonOrderList.OrderItems = orderItems
	return
}

// ------------------ start get -----------------
func (service *SalesPersonService) GetOrderOrderItems(orderId, orderItemId int) (orderItems *models.OrderItem, appErr *errors.AppError) {
	orderItems, err := service.OrderItemRepo.GetOrderItem(orderItemId)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "order not found", Code: -1}
		return
	}
	if orderItems.OrderId != orderId {
		return nil, nil
	}
	products := service.GetAllOrderItemProducts(orderItemId)
	orderItems.Products = products
	return
}

// ------------------ start get -----------------
func (service *SalesPersonService) GetOrderItemProducts(orderItemId, productId int) (products *models.Product, appErr *errors.AppError) {
	products, err := service.ProductRepo.GetProduct(productId)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "Products not found", Code: -1}
		return
	}
	return
}

// ------------------ start get -----------------
func (service *SalesPersonService) GetAllOrderItemProducts(orderItemId int) (products []models.Product) {
	products = service.ProductRepo.GetAllOrderItemProducts(orderItemId)
	return
}

// ------------------ end get-----------------
// ------------------ start get -----------------
func (service *SalesPersonService) GetAllOrderOrderItems(orderId int) (orderItems []models.OrderItem) {
	orderItems = service.OrderItemRepo.GetAllOrderOrderItems(orderId)
	for _, orderItem := range orderItems {
		orderItemId := orderItem.Id
		products := service.GetAllOrderItemProducts(orderItemId)
		orderItem.Products = products
	}
	return
}

// ------------------ end get-----------------
// ------------------ start get -----------------
func (service *SalesPersonService) GetAllSalesPersonSalesPersonOrderList(salesPersonId int) (salesPersonOrderList []models.Order) {
	salesPersonOrderList = service.OrderRepo.GetAllSalesPersonSalesPersonOrderList(salesPersonId)
	for _, order := range salesPersonOrderList {
		orderId := order.Id
		orderItems := service.GetAllOrderOrderItems(orderId)
		order.OrderItems = orderItems
	}
	return
}

// ------------------ end get-----------------
// ------------------ end get-----------------
// ------------------ end get-----------------
// ------------------ end get-----------------

func (service *SalesPersonService) DeleteSalesPerson(salesPersonId int) (appErr *errors.AppError) {
	salesPerson, appErr := service.GetSalesPerson(salesPersonId)
	if appErr != nil {
		return appErr
	}
	service.DeleteAllSalesPersonSalesPersonOrderList(salesPerson)

	service.SalesPersonRepo.DeleteSalesPerson(salesPersonId)

	return
}
func (service *SalesPersonService) DeleteSalesPersonSalesPersonOrderList(salesPersonId, orderId int) (appErr *errors.AppError) {
	order, appErr := service.GetSalesPersonSalesPersonOrderList(salesPersonId, orderId)
	if appErr != nil {
		return appErr
	}
	service.DeleteAllOrderOrderItems(order)

	service.OrderRepo.DeleteOrder(orderId)

	return
}
func (service *SalesPersonService) DeleteOrderOrderItems(orderId, orderItemId int) (appErr *errors.AppError) {
	orderItem, appErr := service.GetOrderOrderItems(orderId, orderItemId)
	if appErr != nil {
		return appErr
	}
	service.DeleteAllOrderItemProducts(orderItem)
	service.OrderItemRepo.DeleteOrderItem(orderItemId)

	return
}
func (service *SalesPersonService) DeleteOrderItemProducts(orderItemId, productId int) (appErr *errors.AppError) {
	_, appErr = service.GetOrderItemProducts(orderItemId, productId)
	if appErr != nil {
		return appErr
	}
	service.ProductRepo.DeleteProduct(productId)

	return
}
func (service *SalesPersonService) DeleteAllOrderItemProducts(orderItem *models.OrderItem) (appErr *errors.AppError) {
	service.OrderItemRepo.DeleteAllOrderItemProducts(orderItem.Id)
	return
}
func (service *SalesPersonService) DeleteAllOrderOrderItems(order *models.Order) (appErr *errors.AppError) {
	service.deleteOrderItemManyReferences(order.OrderItems)
	service.OrderRepo.DeleteAllOrderOrderItems(order.Id)
	return
}
func (service *SalesPersonService) DeleteAllSalesPersonSalesPersonOrderList(salesPerson *models.SalesPerson) (appErr *errors.AppError) {
	service.deleteOrderManyReferences(salesPerson.SalesPersonOrderList)
	service.SalesPersonRepo.DeleteAllSalesPersonSalesPersonOrderList(salesPerson.Id)
	return
}

// start loop delete method
func (service *SalesPersonService) deleteOrderManyReferences(orders []models.Order) {
	for _, ref := range orders {
		order := &ref
		service.DeleteAllOrderOrderItems(order)

	}
	return
}

// end loop delete method
// start loop delete method
func (service *SalesPersonService) deleteOrderItemManyReferences(orderItems []models.OrderItem) {
	for _, ref := range orderItems {
		orderItem := &ref
		service.DeleteAllOrderItemProducts(orderItem)
	}
	return
}

// end loop delete method

func (service *SalesPersonService) UpdateSalesPerson(salesPerson *models.SalesPerson) (appErr *errors.AppError) {
	existingSalesPerson, err := service.GetSalesPerson(salesPerson.Id)
	if err != nil {
		return err
	}
	salesPersonUpdateErr := service.SalesPersonRepo.UpdateSalesPerson(salesPerson)
	if salesPersonUpdateErr != nil {
		appErr = &errors.AppError{Error: salesPersonUpdateErr, Message: "salesPerson could not be updated", Code: -1}
		return appErr
	}
	service.upsertAllSalesPersonSalesPersonOrderList(existingSalesPerson.Id, existingSalesPerson.SalesPersonOrderList, salesPerson.SalesPersonOrderList)

	return appErr
}

func (service *SalesPersonService) upsertAllSalesPersonSalesPersonOrderList(salesPersonId int, existingSalesPersonOrderList []models.Order, newSalesPersonOrderList []models.Order) {
	ordersToBeDeleted, ordersToBeAdded, ordersToBeUpdated := models.GetOrdersTobeDeletedAndAdded(existingSalesPersonOrderList, newSalesPersonOrderList)

	for _, tu := range ordersToBeUpdated {
		service.UpdateSalesPersonSalesPersonOrderList(salesPersonId, &tu)
	}
	for _, ta := range ordersToBeAdded {
		service.CreateSalesPersonSalesPersonOrderList(salesPersonId, &ta)
	}
	for _, td := range ordersToBeDeleted {
		service.DeleteSalesPersonSalesPersonOrderList(salesPersonId, td.Id)
	}
}

// ------------------ start update -----------------
func (service *SalesPersonService) UpdateSalesPersonSalesPersonOrderList(salesPersonId int, order *models.Order) (appErr *errors.AppError) {
	existingOrder, err := service.GetSalesPersonSalesPersonOrderList(salesPersonId, order.Id)
	if err != nil {
		return err
	}
	orderUpdateErr := service.OrderRepo.UpdateOrder(order)
	if orderUpdateErr != nil {
		appErr = &errors.AppError{Error: orderUpdateErr, Message: "order could not be updated", Code: -1}
		return appErr
	}
	service.upsertAllOrderOrderItems(existingOrder.Id, existingOrder.OrderItems, order.OrderItems)

	return appErr
}

func (service *SalesPersonService) upsertAllOrderOrderItems(orderId int, existingOrderItems []models.OrderItem, newOrderItems []models.OrderItem) {
	orderItemsToBeDeleted, orderItemsToBeAdded, orderItemsToBeUpdated := models.GetOrderItemsTobeDeletedAndAdded(existingOrderItems, newOrderItems)

	for _, tu := range orderItemsToBeUpdated {
		service.UpdateOrderOrderItems(orderId, &tu)
	}
	for _, ta := range orderItemsToBeAdded {
		service.CreateOrderOrderItems(orderId, &ta)
	}
	for _, td := range orderItemsToBeDeleted {
		service.DeleteOrderOrderItems(orderId, td.Id)
	}
}

// ------------------ start update -----------------
func (service *SalesPersonService) UpdateOrderOrderItems(orderId int, orderItem *models.OrderItem) (appErr *errors.AppError) {
	existingOrderItem, err := service.GetOrderOrderItems(orderId, orderItem.Id)
	if err != nil {
		return err
	}
	orderItemUpdateErr := service.OrderItemRepo.UpdateOrderItem(orderItem)
	if orderItemUpdateErr != nil {
		appErr = &errors.AppError{Error: orderItemUpdateErr, Message: "orderItem could not be updated", Code: -1}
		return appErr
	}
	service.upsertAllOrderItemProducts(existingOrderItem.Id, existingOrderItem.Products, orderItem.Products)

	return appErr
}

func (service *SalesPersonService) upsertAllOrderItemProducts(orderItemId int, existingProducts []models.Product, newProducts []models.Product) {
	productsToBeDeleted, productsToBeAdded, productsToBeUpdated := models.GetProductsTobeDeletedAndAdded(existingProducts, newProducts)

	for _, tu := range productsToBeUpdated {
		service.UpdateOrderItemProducts(orderItemId, &tu)
	}
	for _, ta := range productsToBeAdded {
		service.CreateOrderItemProducts(orderItemId, &ta)
	}
	for _, td := range productsToBeDeleted {
		service.OrderItemRepo.UnlinkOrderItemProducts(orderItemId, td.Id)
		service.DeleteOrderItemProducts(orderItemId, td.Id)
	}
}

// ------------------ start update -----------------
func (service *SalesPersonService) UpdateOrderItemProducts(orderItemId int, product *models.Product) (appErr *errors.AppError) {
	_, err := service.GetOrderItemProducts(orderItemId, product.Id)
	if err != nil {
		return err
	}
	productUpdateErr := service.ProductRepo.UpdateProduct(product)
	if productUpdateErr != nil {
		appErr = &errors.AppError{Error: productUpdateErr, Message: "product could not be updated", Code: -1}
		return appErr
	}
	return appErr
}

// ------------------ end update-----------------
// ------------------ end update-----------------
// ------------------ end update-----------------

func (service *SalesPersonService) CreateSalesPerson(salesPerson *models.SalesPerson) (salesPersonId int, appErr *errors.AppError) {
	salesPersonId, err := service.SalesPersonRepo.CreateSalesPerson(salesPerson)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "salesPerson could not be created", Code: -1}
	}
	for _, salesPersonOrderList := range salesPerson.SalesPersonOrderList {
		service.CreateSalesPersonSalesPersonOrderList(salesPersonId, &salesPersonOrderList)
	}
	return salesPersonId, appErr
}

func (service *SalesPersonService) CreateSalesPersonSalesPersonOrderList(salesPersonId int, order *models.Order) (orderId int, appErr *errors.AppError) {
	order.SalesPerson = salesPersonId
	orderId, err := service.OrderRepo.CreateOrder(order)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "salesPerson could not be created", Code: -1}
	}
	for _, orderItems := range order.OrderItems {
		service.CreateOrderOrderItems(salesPersonId, &orderItems)
	}
	return salesPersonId, appErr
}

func (service *SalesPersonService) CreateOrderOrderItems(orderId int, orderItem *models.OrderItem) (orderItemId int, appErr *errors.AppError) {
	orderItem.OrderId = orderId
	orderItemId, err := service.OrderItemRepo.CreateOrderItem(orderItem)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "order could not be created", Code: -1}
	}
	for _, products := range orderItem.Products {
		service.CreateOrderItemProducts(orderId, &products)
	}
	return orderId, appErr
}

func (service *SalesPersonService) CreateOrderItemProducts(orderItemId int, product *models.Product) (productId int, appErr *errors.AppError) {
	productId, err := service.ProductRepo.CreateProduct(product)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "orderItem could not be created", Code: -1}
	}
	service.OrderItemRepo.LinkOrderItemProducts(orderItemId, productId)
	return orderItemId, appErr
}
