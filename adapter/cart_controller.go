package adapter

import (
	"amazon-go/domain"
	"amazon-go/usecase"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CartController struct {
	cartUseCase *usecase.CartUseCase
}

func NewCartController(cartUseCase *usecase.CartUseCase) *CartController {
	return &CartController{cartUseCase: cartUseCase}
}

func (c *CartController) AddToCart(ctx echo.Context) error {
	req := new(addToCartRequest)
	if err := ctx.Bind(req); err != nil {
		return err
	}

	cart := &domain.Cart{
		CartId:    uuid.New().String(),
		UserId:    req.UserId,
		ProductId: req.ProductId,
		Quantity:  req.Quantity,
	}

	if err := c.cartUseCase.AddToCart(cart); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "カートへの追加に失敗しました: " + err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{
		"message": "カートに追加されました",
	})
}

type addToCartRequest struct {
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
