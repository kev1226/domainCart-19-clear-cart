package routes

import (
	"clear-cart/view"

	"github.com/gin-gonic/gin"
	"github.com/kev1226/auth-common-go/jwt"
)

func SetupRoutes(r *gin.Engine) {
	group := r.Group("/cart", jwt.AuthGuard("user"))
	group.DELETE("", view.ClearCart)
}
