package config

import (
	"YellowPepper-FundsTransfers/pkg/app/core/ws"
	"YellowPepper-FundsTransfers/pkg/app/mgmt"
	"YellowPepper-FundsTransfers/pkg/app/misc/environment"

	"github.com/gin-gonic/gin"
)

var singletonEngine *gin.Engine

func StopWebServer() {
	//Nothing to do here yet
}

func InitWebServer(environment environment.Environment, accountWs ws.AccountWs, transferWs ws.TransferWs,
	infoWs mgmt.InfoWs, envWs mgmt.EnvWs, metricsWs mgmt.MetricsWs, healthWs mgmt.HealthWs) error {

	singletonEngine = gin.Default()

	loadApiRoutes(accountWs, transferWs)
	loadMgmtRoutes(infoWs, envWs, metricsWs, healthWs)

	hostAddress := environment.GetValueOrDefault(HOST_POST, HOST_POST_DEFAULT_VALUE).AsString()
	return singletonEngine.Run(hostAddress)
}

func loadApiRoutes(accountWs ws.AccountWs, transferWs ws.TransferWs) {

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

func loadMgmtRoutes(infoWs mgmt.InfoWs, envWs mgmt.EnvWs, metricsWs mgmt.MetricsWs, healthWs mgmt.HealthWs) {

	mgmt := singletonEngine.Group("/mgmt")
	mgmt.GET("/info", infoWs.Get)
	mgmt.GET("/env", envWs.Get)
	mgmt.GET("/metrics", metricsWs.Get)
	mgmt.GET("/health", healthWs.Get)
}
