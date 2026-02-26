package models

import "time"

type Payment struct {
	Base
	OrderId              uint           `gorm:"not null;foreignKey:order_id"`
	Order                Order          `gorm:"foreignKey:order_id"`
	Amount               float64        `gorm:"not null"`
	Status               PaymentStatus  `gorm:"type:varchar(20);not null;default:'pending'"`
	GatewayTransactionId string         `gorm:"type:varchar(100);unique"`
	GatewayResponse      string         `gorm:"type:text"`
	PaymentMethod        PaymentMethod  `gorm:"type:varchar(20);not null;default:'credit_card'"`
	PaymentGateway       PaymentGateway `gorm:"type:varchar(20);not null;default:'stripe'"`
	Currency             string         `gorm:"type:varchar(10);not null;default:'USD'"`
	PaidAt               *time.Time     `gorm:"type:timestamp"`
}

func (p *Payment) TableName() string {
	return "payments"
}

type PaymentStatus string

const (
	PaymentStatusPending           PaymentStatus = "pending"
	PaymentStatusCompleted         PaymentStatus = "completed"
	PaymentStatusAuthorized        PaymentStatus = "authorized"
	PaymentStatusCaptured          PaymentStatus = "captured"
	PaymentStatusFailed            PaymentStatus = "failed"
	PaymentStatusRefunded          PaymentStatus = "refunded"
	PaymentStatusPartiallyRefunded PaymentStatus = "partially_refunded"
)

type PaymentMethod string

const (
	PaymentMethodCreditCard   PaymentMethod = "credit_card"
	PaymentMethodDebitCard    PaymentMethod = "debit_card"
	PaymentMethodPayPal       PaymentMethod = "paypal"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
)

type PaymentGateway string

const (
	PaymentGatewayStripe       PaymentGateway = "stripe"
	PaymentGatewayPayPal       PaymentGateway = "paypal"
	PaymentGatewaySquare       PaymentGateway = "square"
	PaymentGatewayAuthorizeNet PaymentGateway = "authorize_net"
)
