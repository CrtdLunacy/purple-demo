package order

import (
	"go/order-api/internal/product"
)

type OrderService struct {
	OrderRepository *OrderRepository
}

type OrderData struct {
	UserId     uint
	Products   []uint
	TotalPrice float64
	UserPhone  string
}

func NewOrderService(orderRepository *OrderRepository) *OrderService {
	return &OrderService{
		OrderRepository: orderRepository,
	}
}

func (service *OrderService) CreateOrder(data *OrderData) (*Order, error) {
	var products []product.Product

	err := service.OrderRepository.Database.DB.
		Where("id IN ?", data.Products).
		Find(&products).Error

	if err != nil {
		return nil, err
	}

	if len(products) != len(data.Products) {
		return nil, err
	}

	newOrder := NewOrder(
		Order{
			UserID:    data.UserId,
			Products:  products,
			Price:     data.TotalPrice,
			UserPhone: data.UserPhone,
		})

	createdOrder, err := service.OrderRepository.Create(newOrder)
	if err != nil {
		return nil, err
	}

	return createdOrder, nil
}
