package config

import (
	coreWs "YellowPepper-FundsTransfers/pkg/core/ws"
	mgmtWs "YellowPepper-FundsTransfers/pkg/mgmt/ws"

	"github.com/gin-gonic/gin"
)

func LoadApiRoutes(router *gin.Engine, accountWs coreWs.AccountWs, transferWs coreWs.TransferWs) {

	api := router.Group("/api")

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

func LoadMgmtRoutes(router *gin.Engine, infoWs mgmtWs.InfoWs, envWs mgmtWs.EnvWs, metricsWs mgmtWs.MetricsWs, healthWs mgmtWs.HealthWs) {
	mgmt := router.Group("/mgmt")
	mgmt.GET("/info", infoWs.Get)
	mgmt.GET("/env", envWs.Get)
	mgmt.GET("/metrics", metricsWs.Get)
	mgmt.GET("/health", healthWs.Get)
}
