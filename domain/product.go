package domain

type Product struct {
	ProductId   string
	Name        string
	Description string
	Price       float64
}

type ProductRepository interface {
	GetProducts() ([]Product, error)
}
