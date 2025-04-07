package usecase

import (
	"amazon-go/domain"

	"github.com/google/uuid"
)

type UserUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(userRepo domain.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (uc *UserUseCase) CreateUser(name, phoneNumber, gender string) (string, error) {
	id, _ := uuid.NewUUID()
	user := &domain.User{
		UserId:      id.String(),
		UserName:    name,
		PhoneNumber: phoneNumber,
		Gender:      gender,
	}

	err := uc.userRepo.Create(user)
	if err != nil {
		return "", err
	}
	return user.UserId, nil
}

type ProductUseCase struct {
	productRepo domain.ProductRepository
}

func NewProductUseCase(productRepo domain.ProductRepository) *ProductUseCase {
	return &ProductUseCase{productRepo: productRepo}
}

func (uc *ProductUseCase) GetProducts() ([]domain.Product, error) {
	return uc.productRepo.GetProducts()
}

type CartUseCase struct {
	cartRepo domain.CartRepository
}

func NewCartUseCase(cartRepo domain.CartRepository) *CartUseCase {
	return &CartUseCase{cartRepo: cartRepo}
}

func (uc *CartUseCase) AddToCart(cart *domain.Cart) error {
	return uc.cartRepo.AddToCart(cart)
}
