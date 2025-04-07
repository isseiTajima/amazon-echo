package usecase

import "amazon-go/domain"

type ProductUseCase struct {
	productRepo domain.ProductRepository
}

func NewProductUseCase(productRepo domain.ProductRepository) *ProductUseCase {
	return &ProductUseCase{productRepo: productRepo}
}

func (uc *ProductUseCase) GetProducts() ([]domain.Product, error) {
	return uc.productRepo.GetProducts()
}
