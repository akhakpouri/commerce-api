package order

import (
	orderitem "commerce/api/internal/dto/order-item"
	"commerce/internal/shared/models"
)

type Order struct {
	Id          uint                  `json:"id"`
	UserId      uint                  `json:"user_id"`
	TotalAmount float64               `json:"total_amount"`
	Status      string                `json:"status"`
	OrderItems  []orderitem.OrderItem `json:"order_items,omitempty"`
}

func FromModel(order *models.Order) *Order {
	orderItems := make([]orderitem.OrderItem, len(order.OrderItems))
	for i, item := range order.OrderItems {
		orderItems[i] = *orderitem.FromModel(&item)
	}

	return &Order{
		Id:          order.Id,
		UserId:      order.UserId,
		TotalAmount: order.TotalAmount,
		Status:      string(order.Status),
		OrderItems:  orderItems,
	}
}

func ToModel(order *Order) *models.Order {
	orderItems := make([]models.OrderItem, len(order.OrderItems))
	for i, item := range order.OrderItems {
		orderItems[i] = *orderitem.ToModel(&item)
	}

	return &models.Order{
		UserId:      order.UserId,
		TotalAmount: order.TotalAmount,
		Status:      models.OrderStatus(order.Status),
		OrderItems:  orderItems,
	}
}
