package application

import (
	"YellowPepper-FundsTransfers/config"
	"YellowPepper-FundsTransfers/core/repository"
	"YellowPepper-FundsTransfers/core/service"
	"YellowPepper-FundsTransfers/core/ws"
	"YellowPepper-FundsTransfers/misc/collection"
	datasource "YellowPepper-FundsTransfers/misc/datasource/impl"
	"YellowPepper-FundsTransfers/misc/transaction"
)

func wire(applicationProperties *collection.Properties) (ws.AccountWs, ws.TransferWs) {

	accountRepository := repository.NewDefaultAccountRepository()
	transferRepository := repository.NewDefaultTransferRepository()

	FUND_TRANSFERS_DATASOURCE_USERNAME := applicationProperties.Get("FUND_TRANSFERS_DATASOURCE_USERNAME")
	FUND_TRANSFERS_DATASOURCE_PASSWORD := applicationProperties.Get("FUND_TRANSFERS_DATASOURCE_PASSWORD")
	FUND_TRANSFERS_DATASOURCE_URL := applicationProperties.Get("FUND_TRANSFERS_DATASOURCE_URL")

	dataSource := datasource.NewMysqlDataSource(FUND_TRANSFERS_DATASOURCE_USERNAME, FUND_TRANSFERS_DATASOURCE_PASSWORD, FUND_TRANSFERS_DATASOURCE_URL)
	transactionHandler := transaction.NewDefaultDBTransactionHandler(dataSource)

	accountService := service.NewDefaultAccountService(transactionHandler, accountRepository)
	transferService := service.NewDefaultTransferService(transactionHandler, transferRepository, accountRepository)

	accountWs := ws.NewDefaultAccountWs(accountService)
	transferWs := ws.NewDefaultTransferWs(transferService)

	return accountWs, transferWs
}

func Run(args []string) {

	applicationProperties := config.LoadApplicationProperties(args)
	accountWs, transferWs := wire(applicationProperties)
	router := config.LoadApiRoutes(accountWs, transferWs)
	router.Run(":8080")
}
