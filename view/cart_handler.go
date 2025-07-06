package view

import (
	"clear-cart/presenter"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ClearCart(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se pudo obtener el usuario"})
		return
	}
	userID := fmt.Sprintf("%v", userIDInterface)

	err := presenter.ClearCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo vaciar el carrito"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Carrito vaciado correctamente"})
}
