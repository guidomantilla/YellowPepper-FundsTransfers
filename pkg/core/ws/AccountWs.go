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
		exception := exception.BadRequestException("error unmarshalling request json to object", err)
		context.JSON(exception.Code, exception)
		return
	}

	if exception := ws.accountService.Create(&account); exception != nil {
		context.JSON(exception.Code, exception)
		return
	}

	context.JSON(http.StatusCreated, account)
}

func (ws DefaultAccountWs) Update(context *gin.Context) {

	id, err := strconv.ParseInt(context.Param("id"), 10, 0)
	if err != nil {
		exception := exception.BadRequestException("url path has an invalid id", err)
		context.JSON(exception.Code, exception)
		return
	}

	var account model.Account
	if err := context.ShouldBindJSON(&account); err != nil {
		exception := exception.BadRequestException("error unmarshalling request json to object", err)
		context.JSON(exception.Code, exception)
		return
	}

	if id != account.Id {
		exception := exception.BadRequestException("url path has an invalid id", err)
		context.JSON(exception.Code, exception)
		return
	}

	if exception := ws.accountService.Update(&account); exception != nil {
		context.JSON(exception.Code, exception)
		return
	}

	context.JSON(http.StatusOK, account)
}

func (ws DefaultAccountWs) Delete(context *gin.Context) {

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

	exception := ws.accountService.DeleteById(id)
	if exception != nil {
		context.JSON(exception.Code, exception)
		return
	}

	context.Status(http.StatusOK)
}

func (ws DefaultAccountWs) FindById(context *gin.Context) {

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

	account, exception := ws.accountService.FindById(id)
	if exception != nil {
		context.JSON(exception.Code, exception)
		return
	}

	context.JSON(http.StatusOK, account)
}

func (ws DefaultAccountWs) FindAll(context *gin.Context) {

	status := context.Query("status")
	fmt.Println(status)

	accounts, exception := ws.accountService.FindAll()
	if exception != nil {
		context.JSON(exception.Code, exception)
		return
	}

	context.JSON(http.StatusOK, accounts)
}
