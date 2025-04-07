package usecase

import "amazon-go/domain"

type CartUseCase struct {
	cartRepo domain.CartRepository
}

func NewCartUseCase(cartRepo domain.CartRepository) *CartUseCase {
	return &CartUseCase{cartRepo: cartRepo}
}

func (uc *CartUseCase) AddToCart(cart *domain.Cart) error {
	return uc.cartRepo.AddToCart(cart)
}
