package main

import (
	"search-address-service/internal/presentation/web_api/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(addressHandlers *handlers.AddressHandlers) *gin.Engine {
	router := gin.Default()

	router.GET("/address/:cep", addressHandlers.GetAddress)

	return router
}
