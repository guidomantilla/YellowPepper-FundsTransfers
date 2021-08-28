package config

import (
	"YellowPepper-FundsTransfers/core/ws"

	"github.com/gin-gonic/gin"
)

var singletonRouter *gin.Engine

func LoadApiRoutes(accountWs ws.AccountWs, transferWs ws.TransferWs) *gin.Engine {
	if singletonRouter == nil {
		loadApiRoutes(accountWs, transferWs)
	}
	return singletonRouter
}

func loadApiRoutes(accountWs ws.AccountWs, transferWs ws.TransferWs) {
	singletonRouter = gin.Default()
	api := singletonRouter.Group("/api")
	{
		api.GET("/transfers", transferWs.FindTransfers)
		api.GET("/transfers/:id", transferWs.FindTransfer)

		api.POST("/transfers", transferWs.DoTransfer)

		api.GET("/accounts", accountWs.FindAll)
		api.GET("/accounts/:id", accountWs.FindById)

		api.POST("/accounts", accountWs.Create)

		api.PUT("/accounts", accountWs.Update)
		api.PUT("/accounts/:id", accountWs.Update)

		api.DELETE("/accounts", accountWs.Delete)
		api.DELETE("/accounts/:id", accountWs.Delete)
	}
}
