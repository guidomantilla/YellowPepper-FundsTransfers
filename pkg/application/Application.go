package application

import (
	"YellowPepper-FundsTransfers/pkg/application/config"
	"YellowPepper-FundsTransfers/pkg/core/repository"
	"YellowPepper-FundsTransfers/pkg/core/service"
	coreWs "YellowPepper-FundsTransfers/pkg/core/ws"
	mgmtWs "YellowPepper-FundsTransfers/pkg/mgmt/ws"
	datasource "YellowPepper-FundsTransfers/pkg/misc/datasource/impl"
	"YellowPepper-FundsTransfers/pkg/misc/environment"
	"YellowPepper-FundsTransfers/pkg/misc/transaction"
)

const (
	FUND_TRANSFERS_DATASOURCE_USERNAME = "FUND_TRANSFERS_DATASOURCE_USERNAME"
	FUND_TRANSFERS_DATASOURCE_PASSWORD = "FUND_TRANSFERS_DATASOURCE_PASSWORD"
	FUND_TRANSFERS_DATASOURCE_URL      = "FUND_TRANSFERS_DATASOURCE_URL"
	HOST_POST                          = "HOST_POST"
	HOST_POST_DEFAULT_VALUE            = ":8080"
	PROFILE                            = "PROFILE"
	PROFILE_DEFAULT_VALUE              = "default"
)

func Run(args *[]string) error {

	env := config.LoadEnvironment(args)
	accountWs, transferWs := wireApi(env)
	infoWs, envWs, metricsWs, healthWs := wireMgmt(env)

	router := config.NewDefaultApiConfig().LoadApiRoutes(accountWs, transferWs).LoadMgmtRoutes(infoWs, envWs, metricsWs, healthWs).GetEngine()
	hostAddress := env.GetValueOrDefault(HOST_POST, HOST_POST_DEFAULT_VALUE).AsString()
	return router.Run(hostAddress)
}

func wireApi(environment environment.Environment) (coreWs.AccountWs, coreWs.TransferWs) {

	FundTransfersDatasourceUsername := environment.GetValue(FUND_TRANSFERS_DATASOURCE_USERNAME).AsString()
	FundTransfersDatasourcePassword := environment.GetValue(FUND_TRANSFERS_DATASOURCE_PASSWORD).AsString()
	FundTransfersDatasourceUrl := environment.GetValue(FUND_TRANSFERS_DATASOURCE_URL).AsString()

	dataSource := datasource.NewMysqlDataSource(FundTransfersDatasourceUsername, FundTransfersDatasourcePassword, FundTransfersDatasourceUrl)
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
