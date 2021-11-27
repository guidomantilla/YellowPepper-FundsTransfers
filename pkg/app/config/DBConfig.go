package config

import (
	"YellowPepper-FundsTransfers/pkg/app/misc/datasource"
	"YellowPepper-FundsTransfers/pkg/app/misc/environment"

	"go.uber.org/zap"
)

var singletonDataSource datasource.DBDataSource

func StopDB() {

	if err := singletonDataSource.GetDatabase().Close(); err != nil {
		zap.L().Fatal("Error closing DB")
		return
	}
}

func InitDB(environment environment.Environment) datasource.DBDataSource {
	username := environment.GetValue(FUND_TRANSFERS_DATASOURCE_USERNAME).AsString()
	password := environment.GetValue(FUND_TRANSFERS_DATASOURCE_PASSWORD).AsString()
	url := environment.GetValue(FUND_TRANSFERS_DATASOURCE_URL).AsString()
	singletonDataSource = datasource.NewMysqlDataSource(username, password, url)
	return singletonDataSource
}
