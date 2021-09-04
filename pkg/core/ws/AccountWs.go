package ws

import (
	"YellowPepper-FundsTransfers/pkg/core/exception"
	"YellowPepper-FundsTransfers/pkg/core/model"
	"YellowPepper-FundsTransfers/pkg/core/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/* TYPES DEFINITION */

type AccountWs interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	FindById(context *gin.Context)
	FindAll(context *gin.Context)
}

type DefaultAccountWs struct {
	accountService service.AccountService
}

/* TYPES CONSTRUCTOR */

func NewDefaultAccountWs(accountService service.AccountService) *DefaultAccountWs {
	return &DefaultAccountWs{
		accountService: accountService,
	}
}

/* DefaultAccountWs METHODS */

func (ws DefaultAccountWs) Create(context *gin.Context) {

	var account model.Account
	if err := context.ShouldBindJSON(&account); err != nil {
		ex := exception.BadRequestException("error unmarshalling request json to object", err)
		context.JSON(ex.Code, ex)
		return
	}

	if ex := ws.accountService.Create(context, &account); ex != nil {
		context.JSON(ex.Code, ex)
		return
	}

	context.JSON(http.StatusCreated, account)
}

func (ws DefaultAccountWs) Update(context *gin.Context) {
	context.Request.Context()
	id, err := strconv.ParseInt(context.Param("id"), 10, 0)
	if err != nil {
		ex := exception.BadRequestException("url path has an invalid id", err)
		context.JSON(ex.Code, ex)
		return
	}

	var account model.Account
	if err := context.ShouldBindJSON(&account); err != nil {
		ex := exception.BadRequestException("error unmarshalling request json to object", err)
		context.JSON(ex.Code, ex)
		return
	}

	if id != account.Id {
		ex := exception.BadRequestException("url path has an invalid id", err)
		context.JSON(ex.Code, ex)
		return
	}

	if ex := ws.accountService.Update(context, &account); ex != nil {
		context.JSON(ex.Code, ex)
		return
	}

	context.JSON(http.StatusOK, account)
}

func (ws DefaultAccountWs) Delete(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 0)
	if err != nil {
		ex := exception.BadRequestException("url path has an invalid id", err)
		context.JSON(ex.Code, ex)
		return
	}

	if context.Request.Body != http.NoBody {
		ex := exception.BadRequestException("body not allowed", nil)
		context.JSON(ex.Code, ex)
		return
	}

	if ex := ws.accountService.DeleteById(context, id); ex != nil {
		context.JSON(ex.Code, ex)
		return
	}

	context.Status(http.StatusOK)
}

func (ws DefaultAccountWs) FindById(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 0)
	if err != nil {
		ex := exception.BadRequestException("url path has an invalid id", err)
		context.JSON(ex.Code, ex)
		return
	}

	if context.Request.Body != http.NoBody {
		ex := exception.BadRequestException("body not allowed", nil)
		context.JSON(ex.Code, ex)
		return
	}

	account, ex := ws.accountService.FindById(context, id)
	if ex != nil {
		context.JSON(ex.Code, ex)
		return
	}

	context.JSON(http.StatusOK, account)
}

func (ws DefaultAccountWs) FindAll(context *gin.Context) {

	status := context.Query("status")
	fmt.Println(status)

	accounts, ex := ws.accountService.FindAll(context)
	if ex != nil {
		context.JSON(ex.Code, ex)
		return
	}

	context.JSON(http.StatusOK, accounts)
}
