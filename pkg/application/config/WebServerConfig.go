package config

import (
	coreWs "YellowPepper-FundsTransfers/pkg/core/ws"
	mgmtWs "YellowPepper-FundsTransfers/pkg/mgmt/ws"
	"YellowPepper-FundsTransfers/pkg/misc/environment"

	"github.com/gin-gonic/gin"
)

var singletonEngine *gin.Engine

func StopWebServer() error {
	return nil
}

func InitWebServer(environment environment.Environment, accountWs coreWs.AccountWs, transferWs coreWs.TransferWs,
	infoWs mgmtWs.InfoWs, envWs mgmtWs.EnvWs, metricsWs mgmtWs.MetricsWs, healthWs mgmtWs.HealthWs) error {

	singletonEngine = gin.Default()

	loadApiRoutes(accountWs, transferWs)
	loadMgmtRoutes(infoWs, envWs, metricsWs, healthWs)

	hostAddress := environment.GetValueOrDefault(HOST_POST, HOST_POST_DEFAULT_VALUE).AsString()
	return singletonEngine.Run(hostAddress)
}

func loadApiRoutes(accountWs coreWs.AccountWs, transferWs coreWs.TransferWs) {

	api := singletonEngine.Group("/api")

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

func loadMgmtRoutes(infoWs mgmtWs.InfoWs, envWs mgmtWs.EnvWs, metricsWs mgmtWs.MetricsWs, healthWs mgmtWs.HealthWs) {

	mgmt := singletonEngine.Group("/mgmt")
	mgmt.GET("/info", infoWs.Get)
	mgmt.GET("/env", envWs.Get)
	mgmt.GET("/metrics", metricsWs.Get)
	mgmt.GET("/health", healthWs.Get)
}
