package service

import (
	"YellowPepper-FundsTransfers/core/exception"
	"YellowPepper-FundsTransfers/core/model"
	"YellowPepper-FundsTransfers/core/repository"
	"YellowPepper-FundsTransfers/core/service/dto"
	"YellowPepper-FundsTransfers/misc/transaction"
	"database/sql"
	"errors"
)

/* TYPES DEFINITION */

type TransferService interface {
	DoTransfer(transferRequest *dto.Transfer) *dto.Transfer
	FindTransfer(id int64) (*model.Transfer, *exception.Exception)
	FindTransfers() (*[]model.Transfer, *exception.Exception)
}

type DefaultTransferService struct {
	transaction.DBTransactionHandler
	transferRepository repository.TransferRepository
	accountRepository  repository.AccountRepository
}

/* TYPES CONSTRUCTOR */

func NewDefaultTransferService(dbTransactionHandler transaction.DBTransactionHandler, transferRepository repository.TransferRepository, accountRepository repository.AccountRepository) *DefaultTransferService {
	return &DefaultTransferService{
		DBTransactionHandler: dbTransactionHandler,
		transferRepository:   transferRepository,
		accountRepository:    accountRepository,
	}
}

/* CONST FOR METHODS */

const (
	EXECUTE_TRANSFER_ERROR_TITLE    = "error executing the transfer"
	FIND_TRANSFER_BY_ID_ERROR_TITLE = "error finding the transfer"
	FIND_TRANSFER_ERROR_TITLE       = "error finding the transfers"
)

/* DefaultTransferService METHODS */

func (service DefaultTransferService) DoTransfer(transferRequest *dto.Transfer) *dto.Transfer {

	transferResponse := &dto.Transfer{
		OriginAccount:      transferRequest.OriginAccount,
		DestinationAccount: transferRequest.DestinationAccount,
		Amount:             transferRequest.Amount,
	}

	if err := createTransferValidation(transferRequest); err != nil {
		transferResponse.Status = "ERROR"
		transferResponse.Errors = *exception.BadRequestException(EXECUTE_TRANSFER_ERROR_TITLE, err)
		return transferResponse
	}

	err := service.HandleTransaction(func(tx *sql.Tx) error {

		originAccount, err := service.accountRepository.FindByNumber(transferRequest.OriginAccount, tx)
		if err != nil {
			return err
		}

		destinationAccount, err := service.accountRepository.FindByNumber(transferRequest.DestinationAccount, tx)
		if err != nil {
			return err
		}

		if originAccount.Balance-transferRequest.Amount < 0 {
			return errors.New("origin account can't have a zero balance")
		}

		originAccount.Balance -= transferRequest.Amount
		destinationAccount.Balance += transferRequest.Amount

		transfer := &model.Transfer{
			OriginAccount: transferRequest.OriginAccount,
			DestinationAccount: transferRequest.DestinationAccount,
			Amount: transferRequest.Amount,
			Date: "DATE",
			Status: "OK",
		}

		if err = service.accountRepository.Update(originAccount, tx); err != nil {
			return err
		}

		if err = service.accountRepository.Update(destinationAccount, tx); err != nil {
			return err
		}

		if err = service.transferRepository.Create(transfer, tx); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		transferResponse.Status = "ERROR"
		transferResponse.Errors = *exception.BadRequestException(EXECUTE_TRANSFER_ERROR_TITLE, err)
		return transferResponse
	}

	transferResponse.Status = "OK"
	return transferResponse
}

func (service DefaultTransferService) FindTransfer(id int64) (*model.Transfer, *exception.Exception) {
	var err error
	var transfer *model.Transfer
	err = service.HandleTransaction(func(tx *sql.Tx) error {

		transfer, err = service.transferRepository.FindById(id, tx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, exception.BadRequestException(FIND_TRANSFER_BY_ID_ERROR_TITLE, err)
	}

	return transfer, nil
}

func (service DefaultTransferService) FindTransfers() (*[]model.Transfer, *exception.Exception) {
	var err error
	var transfer *[]model.Transfer
	err = service.HandleTransaction(func(tx *sql.Tx) error {

		transfer, err = service.transferRepository.FindAll(tx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, exception.BadRequestException(FIND_TRANSFER_ERROR_TITLE, err)
	}

	return transfer, nil
}

/* Helper METHODS */

func createTransferValidation(transferRequest *dto.Transfer) error {
	if transferRequest.OriginAccount < 0 {
		return errors.New("origin account number must be defined")
	}

	if transferRequest.DestinationAccount < 0 {
		return errors.New("destination account number must be defined")
	}

	if transferRequest.Amount < 0 {
		return errors.New("amount must be a positive value")
	}

	return nil
}
