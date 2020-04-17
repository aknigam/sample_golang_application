package service

import (
	"sample_golang_application/errors"
	"sample_golang_application/models"
	"sample_golang_application/repository"
)

type OrderService struct {
	OrderRepo     repository.OrderRepository
	ProductRepo   repository.ProductRepository
	OrderItemRepo repository.OrderItemRepository
}

func (service *OrderService) GetOrder(orderId int) (order *models.Order, appErr *errors.AppError) {
	order, err := service.OrderRepo.GetOrder(orderId)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "order not found", Code: -1}
		return
	}
	orderItems := service.GetAllOrderOrderItems(orderId)
	order.OrderItems = orderItems
	return
}

// ------------------ start get -----------------
func (service *OrderService) GetOrderOrderItems(orderId, orderItemId int) (orderItems *models.OrderItem, appErr *errors.AppError) {
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
func (service *OrderService) GetOrderItemProducts(orderItemId, productId int) (products *models.Product, appErr *errors.AppError) {
	products, err := service.ProductRepo.GetProduct(productId)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "Products not found", Code: -1}
		return
	}
	return
}

// ------------------ start get -----------------
func (service *OrderService) GetAllOrderItemProducts(orderItemId int) (products []models.Product) {
	products = service.ProductRepo.GetAllOrderItemProducts(orderItemId)
	return
}

// ------------------ end get-----------------
// ------------------ start get -----------------
func (service *OrderService) GetAllOrderOrderItems(orderId int) (orderItems []models.OrderItem) {
	orderItems = service.OrderItemRepo.GetAllOrderOrderItems(orderId)
	for _, orderItem := range orderItems {
		orderItemId := orderItem.Id
		products := service.GetAllOrderItemProducts(orderItemId)
		orderItem.Products = products
	}
	return
}

// ------------------ end get-----------------
// ------------------ end get-----------------
// ------------------ end get-----------------

func (service *OrderService) DeleteOrder(orderId int) (appErr *errors.AppError) {
	order, appErr := service.GetOrder(orderId)
	if appErr != nil {
		return appErr
	}
	service.DeleteAllOrderOrderItems(order)

	service.OrderRepo.DeleteOrder(orderId)

	return
}
func (service *OrderService) DeleteOrderOrderItems(orderId, orderItemId int) (appErr *errors.AppError) {
	orderItem, appErr := service.GetOrderOrderItems(orderId, orderItemId)
	if appErr != nil {
		return appErr
	}
	service.DeleteAllOrderItemProducts(orderItem)
	service.OrderItemRepo.DeleteOrderItem(orderItemId)

	return
}
func (service *OrderService) DeleteOrderItemProducts(orderItemId, productId int) (appErr *errors.AppError) {
	_, appErr = service.GetOrderItemProducts(orderItemId, productId)
	if appErr != nil {
		return appErr
	}
	service.ProductRepo.DeleteProduct(productId)

	return
}
func (service *OrderService) DeleteAllOrderItemProducts(orderItem *models.OrderItem) (appErr *errors.AppError) {
	service.OrderItemRepo.DeleteAllOrderItemProducts(orderItem.Id)
	return
}
func (service *OrderService) DeleteAllOrderOrderItems(order *models.Order) (appErr *errors.AppError) {
	service.deleteOrderItemManyReferences(order.OrderItems)
	service.OrderRepo.DeleteAllOrderOrderItems(order.Id)
	return
}

// start loop delete method
func (service *OrderService) deleteOrderItemManyReferences(orderItems []models.OrderItem) {
	for _, ref := range orderItems {
		orderItem := &ref
		service.DeleteAllOrderItemProducts(orderItem)
	}
	return
}

// end loop delete method

func (service *OrderService) UpdateOrder(order *models.Order) (appErr *errors.AppError) {
	existingOrder, err := service.GetOrder(order.Id)
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

func (service *OrderService) upsertAllOrderOrderItems(orderId int, existingOrderItems []models.OrderItem, newOrderItems []models.OrderItem) {
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
func (service *OrderService) UpdateOrderOrderItems(orderId int, orderItem *models.OrderItem) (appErr *errors.AppError) {
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

func (service *OrderService) upsertAllOrderItemProducts(orderItemId int, existingProducts []models.Product, newProducts []models.Product) {
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
func (service *OrderService) UpdateOrderItemProducts(orderItemId int, product *models.Product) (appErr *errors.AppError) {
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

func (service *OrderService) CreateOrder(order *models.Order) (orderId int, appErr *errors.AppError) {
	orderId, err := service.OrderRepo.CreateOrder(order)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "order could not be created", Code: -1}
	}
	for _, orderItems := range order.OrderItems {
		service.CreateOrderOrderItems(orderId, &orderItems)
	}
	return orderId, appErr
}

func (service *OrderService) CreateOrderOrderItems(orderId int, orderItem *models.OrderItem) (orderItemId int, appErr *errors.AppError) {
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

func (service *OrderService) CreateOrderItemProducts(orderItemId int, product *models.Product) (productId int, appErr *errors.AppError) {
	productId, err := service.ProductRepo.CreateProduct(product)
	if err != nil {
		appErr = &errors.AppError{Error: err, Message: "orderItem could not be created", Code: -1}
	}
	service.OrderItemRepo.LinkOrderItemProducts(orderItemId, productId)
	return orderItemId, appErr
}
