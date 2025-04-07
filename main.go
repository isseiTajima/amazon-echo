package main

import (
	"amazon-go/adapter"
	"amazon-go/infrastructure"
	"amazon-go/usecase"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDB() (*gorm.DB, error) {
	dsn := "host=localhost user=user password=postgres dbname=amazon_ec port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := setupDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(&infrastructure.Product{}, &infrastructure.Cart{})

	userRepo := infrastructure.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userController := adapter.NewUserController(userUseCase)

	productRepo := infrastructure.NewUserRepository(db) // Reusing UserRepository for simplicity
	productUseCase := usecase.NewProductUseCase(productRepo)
	productController := adapter.NewProductController(productUseCase)

	cartRepo := infrastructure.NewUserRepository(db) // Reusing UserRepository for simplicity
	cartUseCase := usecase.NewCartUseCase(cartRepo)
	cartController := adapter.NewCartController(cartUseCase)

	e := echo.New()
	e.POST("/users", userController.CreateUser)
	e.GET("/products", productController.GetProducts)
	e.POST("/cart", cartController.AddToCart)

	log.Fatal(e.Start(":1323"))
}
