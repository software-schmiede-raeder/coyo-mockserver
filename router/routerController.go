package router

import (
	eps "./endpoints"
	"github.com/gin-gonic/gin"
)

func AddRoutes(ginRest *gin.Engine) {
	addUserRoutes(ginRest)
}

func addUserRoutes(ginRest *gin.Engine) {
	eps.AddRoutes(ginRest)
}
