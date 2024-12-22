package order

import (
	"go/order-api/pkg/db"
)

type OrderRepository struct {
	Database *db.Db
}

func NewOrderRepository(database *db.Db) *OrderRepository {
	return &OrderRepository{
		Database: database,
	}
}

func (repo *OrderRepository) Create(order *Order) (*Order, error) {
	result := repo.Database.DB.Create(order)
	if result.Error != nil {
		return nil, result.Error
	}

	return order, nil
}

func (repo *OrderRepository) GetByID(userId uint, orderID uint) (*Order, error) {
	var order Order

	result := repo.Database.DB.Preload("Products").Where("user_id = ?", userId).First(&order, orderID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (repo *OrderRepository) GetAll(userId uint) ([]Order, error) {
	var orders []Order

	result := repo.Database.DB.Preload("Products").Where("user_id = ?", userId).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}
