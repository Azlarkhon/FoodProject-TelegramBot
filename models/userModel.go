package models


type CartItem struct {
	ProductID int
	Food      string
	Quantity  int
	Price     float64
}

type User struct {
	ID           int64
	Address      string
	Cart         []CartItem
	OrderHistory []Order
}