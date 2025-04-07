package infrastructure

import (
	"amazon-go/domain"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (r *ProductRepositoryImpl) GetProducts() ([]domain.Product, error) {
	var products []Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	var domainProducts []domain.Product
	for _, p := range products {
		domainProducts = append(domainProducts, domain.Product{
			ProductId:   p.ProductId,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}
	return domainProducts, nil
}
