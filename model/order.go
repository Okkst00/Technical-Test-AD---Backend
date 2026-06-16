package model

import "time"


type Order struct {

	ID int `json:"id"`
	OrderCode string `json:"order_code"`
	UserID int `json:"user_id"`
	CustomerName string `json:"customer_name"`
	CustomerPhone string `json:"customer_phone"`
	CustomerEmail string `json:"customer_email"`
	ShippingAddress string `json:"shipping_address"`
	PaymentMethodID int `json:"payment_method_id"`
	PaymentOptionID *int `json:"payment_option_id"`
	TotalAmount float64 `json:"total_amount"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Items []OrderItem `json:"items"`
}

type OrderItem struct {

	ID int `json:"id"`
	OrderID int `json:"order_id"`
	ProductID int `json:"product_id"`
	ProductName string `json:"product_name"`
	Price int `json:"price"`
	Qty int `json:"qty"`
	Subtotal int `json:"subtotal"`

}

type CreateOrderRequest struct {
	Order struct {
		CustomerName    string `json:"customer_name"`
		CustomerPhone   string `json:"customer_phone"`
		CustomerEmail   string `json:"customer_email"`
		ShippingAddress string `json:"shipping_address"`
		PaymentMethodID int    `json:"payment_method_id"`
		PaymentOptionID *int   `json:"payment_option_id"`
	} `json:"order"`

	Items []CheckoutItem `json:"items"`
}

type OrderItemRequest struct {
	ProductID int `json:"product_id"`
	Qty int `json:"qty"`

}