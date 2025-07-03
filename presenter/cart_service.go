package presenter

import (
	"clear-cart/model"
)

func ClearCart(userID string) error {
	return model.ClearCart(userID)
}
