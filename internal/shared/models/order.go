package models

type Order struct {
	Base
	UserId            uint        `gorm:"not null;"`
	SubTotalAmount    float64     `gorm:"not null"`
	TaxAmount         float64     `gorm:"not null"`
	TotalAmount       float64     `gorm:"not null"`
	OrderNumber       string      `gorm:"type:varchar(100);not null;unique"`
	Status            OrderStatus `gorm:"type:varchar(20);not null;default:'pending'"`
	User              User        `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	ShippingAddress   Address     `gorm:"foreignKey:ShippingAddressId;constraint:OnDelete:RESTRICT"`
	ShippingAddressId uint        `gorm:"not null"`
	BillingAddress    Address     `gorm:"foreignKey:BillingAddressId;constraint:OnDelete:RESTRICT"`
	BillingAddressId  uint        `gorm:"not null"`
	OrderItems        []OrderItem `gorm:"foreignKey:OrderId;constraint:OnDelete:CASCADE"`
	Payments          []Payment   `gorm:"foreignKey:OrderId;constraint:OnDelete:CASCADE"`
}

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

func (o *Order) TableName() string {
	return "orders"
}
