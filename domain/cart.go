package domain

type Cart struct {
	CartId    string
	UserId    string
	ProductId string
	Quantity  int
}

type CartRepository interface {
	AddToCart(cart *Cart) error
}
