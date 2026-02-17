package models

type Order struct {
	Base
	UserId            uint          `gorm:"not null;foreignKey:user_id"`
	SubTotalAmount    float64       `gorm:"not null"`
	TaxAmount         float64       `gorm:"not null"`
	TotalAmount       float64       `gorm:"not null"`
	OrderNumber       string        `gorm:"type:varchar(100);not null;unique"`
	PaymentStatus     PaymentStatus `gorm:"type:varchar(20);not null;default:'pending'"`
	Status            OrderStatus   `gorm:"type:varchar(20);not null;default:'pending'"`
	User              User          `gorm:"foreignKey:user_id"`
	ShippingAddress   Address       `gorm:"foreignKey:shipping_address_id"`
	ShippingAddressId uint          `gorm:"not null"`
	BillingAddress    Address       `gorm:"foreignKey:billing_address_id"`
	BillingAddressId  uint          `gorm:"not null"`
	OrderItems        []OrderItem   `gorm:"foreignKey:order_id"`
}

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusRefunded  PaymentStatus = "refunded"
)

func (o *Order) TableName() string {
	return "orders"
}
