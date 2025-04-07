package infrastructure

import (
	"amazon-go/domain"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type User struct {
	UserId      string `gorm:"primaryKey;type:uuid"`
	UserName    string `gorm:"size:100;not null"`
	PhoneNumber string `gorm:"size:13;not null"`
	Gender      string `gorm:"size:10;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"` //deleteを行う際にDeletedAtに削除日時が入る
}

type Product struct {
	ProductId   string  `gorm:"primaryKey;type:uuid"`
	Name        string  `gorm:"size:100;not null"`
	Description string  `gorm:"size:255"`
	Price       float64 `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Cart struct {
	CartId    string `gorm:"primaryKey;type:uuid"`
	UserId    string `gorm:"not null"`
	ProductId string `gorm:"not null"`
	Quantity  int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{db: db}
}

func toInfrastructureUser(user *domain.User) *User {
	return &User{
		UserId:      user.UserId,
		UserName:    user.UserName,
		PhoneNumber: user.PhoneNumber,
		Gender:      user.Gender,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   gorm.DeletedAt{},
	}
}

func (r *UserRepository) Create(user *domain.User) error {
	infraUser := toInfrastructureUser(user)
	return r.db.Create(infraUser).Error
}

func (r *UserRepository) GetProducts() ([]Product, error) {
	var products []Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *UserRepository) AddToCart(cart *Cart) error {
	return r.db.Create(cart).Error
}
