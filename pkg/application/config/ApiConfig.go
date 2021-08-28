package config

import (
	coreWs "YellowPepper-FundsTransfers/pkg/core/ws"
	mgmtWs "YellowPepper-FundsTransfers/pkg/mgmt/ws"

	"github.com/gin-gonic/gin"
)

type ApiConfig interface {
	GetEngine() *gin.Engine
	LoadApiRoutes(accountWs coreWs.AccountWs, transferWs coreWs.TransferWs) ApiConfig
	LoadMgmtRoutes(infoWs mgmtWs.InfoWs) ApiConfig
}

type DefaultApiConfig struct {
	singletonRouter *gin.Engine
}

func NewDefaultApiConfig() *DefaultApiConfig {
	return &DefaultApiConfig{
		singletonRouter: gin.Default(),
	}
}

func (config *DefaultApiConfig) GetEngine() *gin.Engine {
	return config.singletonRouter
}

func (config *DefaultApiConfig) LoadApiRoutes(accountWs coreWs.AccountWs, transferWs coreWs.TransferWs) *DefaultApiConfig {

	api := config.singletonRouter.Group("/api")
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
	return config
}

func (config *DefaultApiConfig) LoadMgmtRoutes(infoWs mgmtWs.InfoWs, envWs mgmtWs.EnvWs, metricsWs mgmtWs.MetricsWs, healthWs mgmtWs.HealthWs) *DefaultApiConfig {
	mgmt := config.singletonRouter.Group("/mgmt")
	{
		mgmt.GET("/info", infoWs.Get)
		mgmt.GET("/env", envWs.Get)
		mgmt.GET("/metrics", metricsWs.Get)
		mgmt.GET("/health", healthWs.Get)
	}
	return config
}
