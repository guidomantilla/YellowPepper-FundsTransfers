package cmd

import (
	"YellowPepper-FundsTransfers/pkg/app/config"
	"YellowPepper-FundsTransfers/pkg/app/core/repository"
	"YellowPepper-FundsTransfers/pkg/app/core/service"
	"YellowPepper-FundsTransfers/pkg/app/core/ws"
	"YellowPepper-FundsTransfers/pkg/app/mgmt"
	"YellowPepper-FundsTransfers/pkg/app/misc/transaction"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func CreateServeCmd() *cobra.Command {
	return &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {

			environment := config.InitParams(&args)
			defer config.StopParams()

			config.InitMonitoring(environment)
			defer config.StopMonitoring()

			dataSource := config.InitDB(environment)
			defer config.StopDB()

			infoWs := mgmt.NewDefaultInfoWs()
			envWs := mgmt.NewDefaultEnvWs(environment)
			metricsWs := mgmt.NewDefaultMetricsWs()
			healthWs := mgmt.NewDefaultHealthWs()

			transactionHandler := transaction.NewDefaultDBTransactionHandler(dataSource)

			transferRepository := repository.NewDefaultTransferRepository()
			accountRepository := repository.NewDefaultAccountRepository()

			accountService := service.NewDefaultAccountService(transactionHandler, accountRepository)
			transferService := service.NewDefaultTransferService(transactionHandler, transferRepository, accountRepository)

			accountWs := ws.NewDefaultAccountWs(accountService)
			transferWs := ws.NewDefaultTransferWs(transferService)

			if err := config.InitWebServer(environment, accountWs, transferWs, infoWs, envWs, metricsWs, healthWs); err != nil {
				zap.L().Fatal("error starting the server.")
			}
			defer config.StopWebServer()
		},
	}
}
