package service

import (
	"YellowPepper-FundsTransfers/pkg/app/core/exception"
	"YellowPepper-FundsTransfers/pkg/app/core/model"
	"YellowPepper-FundsTransfers/pkg/app/core/repository"
	"YellowPepper-FundsTransfers/pkg/app/misc/transaction"
	"context"
	"database/sql"
	"errors"
	"strings"
)

/* TYPES DEFINITION */

type AccountService interface {
	Create(context context.Context, account *model.Account) *exception.Exception
	Update(context context.Context, account *model.Account) *exception.Exception
	DeleteById(context context.Context, id int64) *exception.Exception
	FindById(context context.Context, id int64) (*model.Account, *exception.Exception)
	FindAll(context context.Context) (*[]model.Account, *exception.Exception)
}

type DefaultAccountService struct {
	transaction.DBTransactionHandler
	accountRepository repository.AccountRepository
}

/* TYPES CONSTRUCTOR */

func NewDefaultAccountService(dbTransactionHandler transaction.DBTransactionHandler, accountRepository repository.AccountRepository) *DefaultAccountService {
	return &DefaultAccountService{
		DBTransactionHandler: dbTransactionHandler,
		accountRepository:    accountRepository,
	}
}

/* DefaultAccountService METHODS */

func (service *DefaultAccountService) Create(context context.Context, account *model.Account) *exception.Exception {

	if err := createAccountValidation(account); err != nil {
		return exception.BadRequestException("error creating the account", err)
	}

	err := service.HandleTransaction(func(tx *sql.Tx) error {

		account.Status = "CREATED"
		if err := service.accountRepository.Create(context, tx, account); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return exception.InternalServerErrorException("error creating the account", err)
	}

	return nil
}

func (service *DefaultAccountService) Update(context context.Context, account *model.Account) *exception.Exception {

	if err := updateAccountValidation(account); err != nil {
		return exception.BadRequestException("error updating the account", err)
	}

	err := service.HandleTransaction(func(tx *sql.Tx) error {

		_, err := service.accountRepository.FindById(context, tx, account.Id)
		if err != nil {
			return err
		}

		if err = service.accountRepository.Update(context, tx, account); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return exception.InternalServerErrorException("error updating the account", err)
	}
	return nil
}

func (service *DefaultAccountService) DeleteById(context context.Context, id int64) *exception.Exception {

	err := service.HandleTransaction(func(tx *sql.Tx) error {

		_, err := service.accountRepository.FindById(context, tx, id)
		if err != nil {
			return err
		}

		if err = service.accountRepository.DeleteById(context, tx, id); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return exception.InternalServerErrorException("error deleting the account", err)
	}
	return nil

}

func (service *DefaultAccountService) FindById(context context.Context, id int64) (*model.Account, *exception.Exception) {

	var err error
	var account *model.Account
	err = service.HandleTransaction(func(tx *sql.Tx) error {

		account, err = service.accountRepository.FindById(context, tx, id)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, exception.InternalServerErrorException("error finding the account", err)
	}

	return account, nil
}

func (service *DefaultAccountService) FindAll(context context.Context) (*[]model.Account, *exception.Exception) {

	var err error
	var accounts *[]model.Account
	err = service.HandleTransaction(func(tx *sql.Tx) error {

		accounts, err = service.accountRepository.FindAll(context, tx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, exception.InternalServerErrorException("error finding the accounts", err)
	}

	return accounts, nil
}

/* Helper METHODS */

func createAccountValidation(account *model.Account) error {
	if account.Id != 0 {
		return errors.New("account id must be undefined")
	}

	if account.Status != "" || strings.TrimSpace(account.Status) != "" {
		return errors.New("account status must be undefined")
	}

	return nil
}

func updateAccountValidation(account *model.Account) error {
	if account.Id == 0 {
		return errors.New("account id must be defined")
	}

	if account.Number == 0 {
		return errors.New("account number must be defined")
	}

	if account.Balance <= 0 {
		return errors.New("account balance must be positive")
	}

	if account.Owner == "" || strings.TrimSpace(account.Owner) == "" {
		return errors.New("account owner must be defined")
	}

	if account.Status == "" || strings.TrimSpace(account.Status) == "" {
		return errors.New("account status must be defined")
	}

	return nil
}
