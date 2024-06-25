package models

import "time"

type Order struct {
	ID          int           `json:"id"`
	UserID      int           `json:"user_id"`
	TotalAmount float64       `json:"total_amount"`
	Status      string        `json:"status"`
	CreatedAt   time.Time     `json:"created_at"`
	Details     []OrderDetail `json:"details"`
}

type OrderDetail struct {
	ID        int     `json:"id"`
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
