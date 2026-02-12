package models

type OrderItem struct {
	Base
	OrderId   uint    `gorm:"not null;foreignKey:order_id"`
	ProductId uint    `gorm:"not null;foreignKey:product_id"`
	Quantity  int     `gorm:"not null"`
	UnitPrice float64 `gorm:"not null"`
	TaxAmount float64 `gorm:"not null"`
	Order     Order   `gorm:"foreignKey:order_id"`
	Product   Product `gorm:"foreignKey:product_id"`
}

func (oi *OrderItem) TableName() string {
	return "order_items"
}
