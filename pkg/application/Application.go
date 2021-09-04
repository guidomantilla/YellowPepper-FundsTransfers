package application

import (
	"YellowPepper-FundsTransfers/pkg/application/config"
	"YellowPepper-FundsTransfers/pkg/core/repository"
	"YellowPepper-FundsTransfers/pkg/core/service"
	coreWs "YellowPepper-FundsTransfers/pkg/core/ws"
	mgmtWs "YellowPepper-FundsTransfers/pkg/mgmt/ws"
	"YellowPepper-FundsTransfers/pkg/misc/datasource"
	"YellowPepper-FundsTransfers/pkg/misc/environment"
	"YellowPepper-FundsTransfers/pkg/misc/transaction"

	"github.com/gin-gonic/gin"
)

var singletonRouter *gin.Engine
var singletonEnvironment environment.Environment

func Run() error {

	singletonEnvironment = config.LoadEnvironment()
	accountWs, transferWs := wireApi(singletonEnvironment)
	infoWs, envWs, metricsWs, healthWs := wireMgmt(singletonEnvironment)

	singletonRouter = gin.Default()
	config.LoadApiRoutes(singletonRouter, accountWs, transferWs)
	config.LoadMgmtRoutes(singletonRouter, infoWs, envWs, metricsWs, healthWs)
	hostAddress := singletonEnvironment.GetValueOrDefault(config.HOST_POST, config.HOST_POST_DEFAULT_VALUE).AsString()
	return singletonRouter.Run(hostAddress)
}

func wireApi(environment environment.Environment) (coreWs.AccountWs, coreWs.TransferWs) {

	username := environment.GetValue(config.FUND_TRANSFERS_DATASOURCE_USERNAME).AsString()
	password := environment.GetValue(config.FUND_TRANSFERS_DATASOURCE_PASSWORD).AsString()
	url := environment.GetValue(config.FUND_TRANSFERS_DATASOURCE_URL).AsString()

	dataSource := datasource.NewMysqlDataSource(username, password, url)
	transactionHandler := transaction.NewDefaultDBTransactionHandler(dataSource)

	transferRepository := repository.NewDefaultTransferRepository()
	accountRepository := repository.NewDefaultAccountRepository()

	accountService := service.NewDefaultAccountService(transactionHandler, accountRepository)
	transferService := service.NewDefaultTransferService(transactionHandler, transferRepository, accountRepository)

	accountWs := coreWs.NewDefaultAccountWs(accountService)
	transferWs := coreWs.NewDefaultTransferWs(transferService)

	return accountWs, transferWs
}

func wireMgmt(environment environment.Environment) (mgmtWs.InfoWs, mgmtWs.EnvWs, mgmtWs.MetricsWs, mgmtWs.HealthWs) {

	infoWs := mgmtWs.NewDefaultInfoWs()
	envWs := mgmtWs.NewDefaultEnvWs(environment)
	metricsWs := mgmtWs.NewDefaultMetricsWs()
	healthWs := mgmtWs.NewDefaultHealthWs()

	return infoWs, envWs, metricsWs, healthWs
}
