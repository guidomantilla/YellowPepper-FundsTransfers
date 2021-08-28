package application

import (
	"YellowPepper-FundsTransfers/config"
	"YellowPepper-FundsTransfers/core/repository"
	"YellowPepper-FundsTransfers/core/service"
	"YellowPepper-FundsTransfers/core/ws"
	datasource "YellowPepper-FundsTransfers/misc/datasource/impl"
	"YellowPepper-FundsTransfers/misc/environment"
	"YellowPepper-FundsTransfers/misc/transaction"
)

func Run(args *[]string) error {

	environment := config.LoadEnvironment(args)
	accountWs, transferWs := wire(environment)
	router := config.LoadApiRoutes(accountWs, transferWs)
	return router.Run(":8080")
}

func wire(environment environment.Environment) (ws.AccountWs, ws.TransferWs) {

	FUND_TRANSFERS_DATASOURCE_USERNAME := environment.Get("FUND_TRANSFERS_DATASOURCE_USERNAME").AsString()
	FUND_TRANSFERS_DATASOURCE_PASSWORD := environment.Get("FUND_TRANSFERS_DATASOURCE_PASSWORD").AsString()
	FUND_TRANSFERS_DATASOURCE_URL := environment.Get("FUND_TRANSFERS_DATASOURCE_URL").AsString()

	dataSource := datasource.NewMysqlDataSource(FUND_TRANSFERS_DATASOURCE_USERNAME, FUND_TRANSFERS_DATASOURCE_PASSWORD, FUND_TRANSFERS_DATASOURCE_URL)
	transactionHandler := transaction.NewDefaultDBTransactionHandler(dataSource)

	transferRepository := repository.NewDefaultTransferRepository()
	accountRepository := repository.NewDefaultAccountRepository()

	accountService := service.NewDefaultAccountService(transactionHandler, accountRepository)
	transferService := service.NewDefaultTransferService(transactionHandler, transferRepository, accountRepository)

	accountWs := ws.NewDefaultAccountWs(accountService)
	transferWs := ws.NewDefaultTransferWs(transferService)

	return accountWs, transferWs
}
