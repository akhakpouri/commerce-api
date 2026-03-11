package models

type OrderItem struct {
	Base
	OrderId   uint    `gorm:"not null;"`
	ProductId uint    `gorm:"not null;"`
	Quantity  int     `gorm:"not null"`
	UnitPrice float64 `gorm:"not null"`
	Order     Order   `gorm:"foreignKey:OrderId;constraint:OnDelete:CASCADE"`
	Product   Product `gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE"`
}

func (oi *OrderItem) TableName() string {
	return "order_items"
}
