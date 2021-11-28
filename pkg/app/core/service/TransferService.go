package service

import (
	"YellowPepper-FundsTransfers/pkg/app/core/exception"
	"YellowPepper-FundsTransfers/pkg/app/core/model"
	"YellowPepper-FundsTransfers/pkg/app/core/repository"
	"YellowPepper-FundsTransfers/pkg/app/core/service/dto"
	"YellowPepper-FundsTransfers/pkg/app/misc/transaction"
	"context"
	"database/sql"
	"errors"
)

/* TYPES DEFINITION */

type TransferService interface {
	DoTransfer(context context.Context, transferRequest *dto.Transfer) *dto.Transfer
	FindTransfer(context context.Context, id int64) (*model.Transfer, *exception.Exception)
	FindTransfers(context context.Context) (*[]model.Transfer, *exception.Exception)
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

/* DefaultTransferService METHODS */

func (service *DefaultTransferService) DoTransfer(context context.Context, transferRequest *dto.Transfer) *dto.Transfer {

	transferResponse := &dto.Transfer{
		OriginAccount:      transferRequest.OriginAccount,
		DestinationAccount: transferRequest.DestinationAccount,
		Amount:             transferRequest.Amount,
	}

	if err := createTransferValidation(transferRequest); err != nil {
		transferResponse.Status, transferResponse.Errors = "ERROR", *exception.BadRequestException("error executing the transfer", err)
		return transferResponse
	}

	err := service.HandleTransaction(func(tx *sql.Tx) error {

		originAccount, err := service.accountRepository.FindByNumber(context, tx, transferRequest.OriginAccount)
		if err != nil {
			return err
		}

		destinationAccount, err := service.accountRepository.FindByNumber(context, tx, transferRequest.DestinationAccount)
		if err != nil {
			return err
		}

		if originAccount.Balance-transferRequest.Amount < 0 {
			return errors.New("origin account can't have a zero balance")
		}

		originAccount.Balance -= transferRequest.Amount
		destinationAccount.Balance += transferRequest.Amount

		transfer := &model.Transfer{
			OriginAccount:      transferRequest.OriginAccount,
			DestinationAccount: transferRequest.DestinationAccount,
			Amount:             transferRequest.Amount,
			Date:               "DATE",
			Status:             "OK",
		}

		if err = service.accountRepository.Update(context, tx, originAccount); err != nil {
			return err
		}

		if err = service.accountRepository.Update(context, tx, destinationAccount); err != nil {
			return err
		}

		if err = service.transferRepository.Create(context, tx, transfer); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		transferResponse.Status, transferResponse.Errors = "ERROR", *exception.BadRequestException("error executing the transfer", err)
		return transferResponse
	}

	transferResponse.Status = "OK"
	return transferResponse
}

func (service *DefaultTransferService) FindTransfer(context context.Context, id int64) (*model.Transfer, *exception.Exception) {
	var err error
	var transfer *model.Transfer
	err = service.HandleTransaction(func(tx *sql.Tx) error {

		transfer, err = service.transferRepository.FindById(context, tx, id)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, exception.BadRequestException("error finding the transfer", err)
	}

	return transfer, nil
}

func (service *DefaultTransferService) FindTransfers(context context.Context) (*[]model.Transfer, *exception.Exception) {
	var err error
	var transfer *[]model.Transfer
	err = service.HandleTransaction(func(tx *sql.Tx) error {

		transfer, err = service.transferRepository.FindAll(context, tx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, exception.BadRequestException("error finding the transfers", err)
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
