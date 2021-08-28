package application

import (
	"YellowPepper-FundsTransfers/pkg/config"
	"YellowPepper-FundsTransfers/pkg/core/repository"
	"YellowPepper-FundsTransfers/pkg/core/service"
	"YellowPepper-FundsTransfers/pkg/core/ws"
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
)

func Run(args *[]string) error {

	env := config.LoadEnvironment(args)
	accountWs, transferWs := wire(env)
	router := config.LoadApiRoutes(accountWs, transferWs)
	hostAddress := env.GetValueOrDefault(HOST_POST, HOST_POST_DEFAULT_VALUE).AsString()
	return router.Run(hostAddress)
}

func wire(environment environment.Environment) (ws.AccountWs, ws.TransferWs) {

	FundTransfersDatasourceUsername := environment.GetValue(FUND_TRANSFERS_DATASOURCE_USERNAME).AsString()
	FundTransfersDatasourcePassword := environment.GetValue(FUND_TRANSFERS_DATASOURCE_PASSWORD).AsString()
	FundTransfersDatasourceUrl := environment.GetValue(FUND_TRANSFERS_DATASOURCE_URL).AsString()

	dataSource := datasource.NewMysqlDataSource(FundTransfersDatasourceUsername, FundTransfersDatasourcePassword, FundTransfersDatasourceUrl)
	transactionHandler := transaction.NewDefaultDBTransactionHandler(dataSource)

	transferRepository := repository.NewDefaultTransferRepository()
	accountRepository := repository.NewDefaultAccountRepository()

	accountService := service.NewDefaultAccountService(transactionHandler, accountRepository)
	transferService := service.NewDefaultTransferService(transactionHandler, transferRepository, accountRepository)

	accountWs := ws.NewDefaultAccountWs(accountService)
	transferWs := ws.NewDefaultTransferWs(transferService)

	return accountWs, transferWs
}
