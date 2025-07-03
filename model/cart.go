package model

import (
	"clear-cart/config"
	"fmt"
)

func ClearCart(userID string) error {
	key := fmt.Sprintf("cart:%s", userID)
	return config.RedisClient.Del(config.Ctx, key).Err()
}
