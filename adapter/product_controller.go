package adapter

import (
	"amazon-go/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productUseCase *usecase.ProductUseCase
}

func NewProductController(productUseCase *usecase.ProductUseCase) *ProductController {
	return &ProductController{productUseCase: productUseCase}
}

func (c *ProductController) GetProducts(ctx echo.Context) error {
	products, err := c.productUseCase.GetProducts()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "商品の取得に失敗しました: " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, products)
}
