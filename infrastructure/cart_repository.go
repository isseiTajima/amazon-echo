package infrastructure

import (
	"amazon-go/domain"
	"time"

	"gorm.io/gorm"
)

type CartRepositoryImpl struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) domain.CartRepository {
	return &CartRepositoryImpl{db: db}
}

func (r *CartRepositoryImpl) AddToCart(cart *domain.Cart) error {
	infraCart := Cart{
		CartId:    cart.CartId,
		UserId:    cart.UserId,
		ProductId: cart.ProductId,
		Quantity:  cart.Quantity,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return r.db.Create(&infraCart).Error
}
