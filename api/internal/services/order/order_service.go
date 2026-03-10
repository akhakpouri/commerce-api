package order

import (
	dto "commerce/api/internal/dto/order"
	"commerce/internal/shared/models"
	repo "commerce/internal/shared/repositories/order"
	"fmt"
	"log/slog"
)

type OrderServiceI interface {
	GetById(id uint) (*dto.Order, error)
	GetByUserId(userId uint) ([]*dto.Order, error)
	Save(order dto.Order) error
	Delete(id uint, hard bool) error
	UpdateStatus(id uint, status string) error
}

type OrderService struct {
	repo repo.OrderRepositoryI
}

func NewOrderService(repo repo.OrderRepositoryI) OrderServiceI {
	return &OrderService{repo: repo}
}

// Delete implements [OrderServiceI].
func (o *OrderService) Delete(id uint, hard bool) error {
	return o.repo.Delete(id, hard)
}

// GetById implements [OrderServiceI].
func (o *OrderService) GetById(id uint) (*dto.Order, error) {
	model, err := o.repo.GetById(id)
	if err != nil {
		slog.Error("Exception occurred getting order by id.", "id", id, "error", err)
		return nil, err
	}
	return dto.FromModel(model), nil
}

// GetByUserId implements [OrderServiceI].
func (o *OrderService) GetByUserId(userId uint) ([]*dto.Order, error) {
	models, err := o.repo.GetAllByUserId(userId)
	if err != nil {
		slog.Error("Exception occurred getting orders by user", "userId", userId, "error", err)
		return nil, err
	}
	orders := make([]*dto.Order, 0, len(models))
	for _, model := range models {
		orders = append(orders, dto.FromModel(model))
	}
	return orders, nil
}

// Save implements [OrderServiceI].
func (o *OrderService) Save(order dto.Order) error {
	model := dto.ToModel(&order)
	return o.repo.Save(model)
}

// UpdateStatus implements [OrderServiceI].
func (o *OrderService) UpdateStatus(id uint, status string) error {
	if !isOrderStatusValid(status) {
		slog.Error("Order status doesn't exist.", "status", status)
		return fmt.Errorf("invalid order status: %s", status)
	}
	return o.repo.UpdateStatus(id, status)
}

func isOrderStatusValid(status string) bool {
	var validStatuses = map[models.OrderStatus]struct{}{
		models.OrderStatusPending:   {},
		models.OrderStatusDelivered: {},
		models.OrderStatusShipped:   {},
		models.OrderStatusCancelled: {},
	}
	_, ok := validStatuses[models.OrderStatus(status)]
	return ok
}
