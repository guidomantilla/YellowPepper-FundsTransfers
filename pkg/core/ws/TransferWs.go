package ws

import (
	"YellowPepper-FundsTransfers/pkg/core/exception"
	"YellowPepper-FundsTransfers/pkg/core/service"
	"YellowPepper-FundsTransfers/pkg/core/service/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransferWs interface {
	DoTransfer(context *gin.Context)
	FindTransfer(context *gin.Context)
	FindTransfers(context *gin.Context)
}

type DefaultTransferWs struct {
	transferService service.TransferService
}

func NewDefaultTransferWs(transferService service.TransferService) *DefaultTransferWs {
	return &DefaultTransferWs{
		transferService: transferService,
	}
}
func (ws DefaultTransferWs) DoTransfer(context *gin.Context) {
	var transferRequest dto.Transfer
	if err := context.ShouldBindJSON(&transferRequest); err != nil {
		exception := exception.BadRequestException("error unmarshalling request json to object", err)
		context.JSON(exception.Code, exception)
		return
	}

	transferResponse := ws.transferService.DoTransfer(&transferRequest)
	if transferResponse.Status == "ERROR" {
		context.JSON(transferResponse.Errors.Code, transferResponse)
		return
	}

	context.JSON(http.StatusCreated, transferResponse)
}

func (ws DefaultTransferWs) FindTransfer(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 0)
	if err != nil {
		exception := exception.BadRequestException("url path has an invalid id", err)
		context.JSON(exception.Code, exception)
		return
	}

	if context.Request.Body != http.NoBody {
		exception := exception.BadRequestException("body not allowed", nil)
		context.JSON(exception.Code, exception)
		return
	}

	account, exception := ws.transferService.FindTransfer(id)
	if exception != nil {
		context.JSON(exception.Code, exception)
		return
	}

	context.JSON(http.StatusOK, account)
}

func (ws DefaultTransferWs) FindTransfers(context *gin.Context) {

	accounts, exception := ws.transferService.FindTransfers()
	if exception != nil {
		context.JSON(exception.Code, exception)
		return
	}

	context.JSON(http.StatusOK, accounts)
}
