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
	"fmt"

	"go.uber.org/zap"
)

func Stop() error {

	return config.StopMonitoring()
}

func Run() error {

	singletonEnvironment := config.LoadEnvironment()
	config.InitMonitoring(singletonEnvironment)

	for _, source := range singletonEnvironment.GetPropertySources() {
		name := source.AsMap()["name"]
		internalMap := source.AsMap()["value"].(map[string]string)
		for key, value := range internalMap {
			zap.L().Debug(fmt.Sprintf("source name: %s, key: %s, value: %s", name, key, value))
			zap.L().Error(fmt.Sprintf("source name: %s, key: %s, value: %s", name, key, value))
		}
	}

	accountWs, transferWs := wireApi(singletonEnvironment)
	infoWs, envWs, metricsWs, healthWs := wireMgmt(singletonEnvironment)

	return config.InitWebServer(singletonEnvironment, accountWs, transferWs, infoWs, envWs, metricsWs, healthWs)
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
