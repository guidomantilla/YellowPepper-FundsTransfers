package service

import (
	"YellowPepper-FundsTransfers/core/exception"
	"YellowPepper-FundsTransfers/core/model"
	"YellowPepper-FundsTransfers/core/repository"
	"YellowPepper-FundsTransfers/misc/transaction"
	"database/sql"
	"errors"
	"strings"
)

/* TYPES DEFINITION */

type AccountService interface {
	Create(account *model.Account) *exception.Exception
	Update(account *model.Account) *exception.Exception
	DeleteById(id int64) *exception.Exception
	FindById(id int64) (*model.Account, *exception.Exception)
	FindAll() (*[]model.Account, *exception.Exception)
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

/* CONST FOR METHODS */

const (
	CREATE_ACCOUNT_ERROR_TITLE     = "error creating the account"
	UPDATE_ACCOUNT_ERROR_TITLE     = "error updating the account"
	DELETE_ACCOUNT_ERROR_TITLE     = "error deleting the account"
	FIND_ACCOUNT_BY_ID_ERROR_TITLE = "error finding the account"
	FIND_ACCOUNT_ERROR_TITLE       = "error finding the accounts"
)

/* DefaultAccountService METHODS */

func (service DefaultAccountService) Create(account *model.Account) *exception.Exception {

	if err := createAccountValidation(account); err != nil {
		return exception.BadRequestException(CREATE_ACCOUNT_ERROR_TITLE, err)
	}

	err := service.HandleTransaction(func(tx *sql.Tx) error {

		account.Status = "CREATED"
		if err := service.accountRepository.Create(account, tx); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return exception.InternalServerErrorException(CREATE_ACCOUNT_ERROR_TITLE, err)
	}

	return nil
}

func (service DefaultAccountService) Update(account *model.Account) *exception.Exception {

	if err := updateAccountValidation(account); err != nil {
		return exception.BadRequestException(UPDATE_ACCOUNT_ERROR_TITLE, err)
	}

	err := service.HandleTransaction(func(tx *sql.Tx) error {

		_, err := service.accountRepository.FindById(account.Id, tx)
		if err != nil {
			return err
		}

		if err = service.accountRepository.Update(account, tx); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return exception.InternalServerErrorException(UPDATE_ACCOUNT_ERROR_TITLE, err)
	}
	return nil
}

func (service DefaultAccountService) DeleteById(id int64) *exception.Exception {

	err := service.HandleTransaction(func(tx *sql.Tx) error {

		_, err := service.accountRepository.FindById(id, tx)
		if err != nil {
			return err
		}

		if err = service.accountRepository.DeleteById(id, tx); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return exception.InternalServerErrorException(DELETE_ACCOUNT_ERROR_TITLE, err)
	}
	return nil

}

func (service DefaultAccountService) FindById(id int64) (*model.Account, *exception.Exception) {

	var err error
	var account *model.Account
	err = service.HandleTransaction(func(tx *sql.Tx) error {

		account, err = service.accountRepository.FindById(id, tx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, exception.InternalServerErrorException(FIND_ACCOUNT_BY_ID_ERROR_TITLE, err)
	}

	return account, nil
}

func (service DefaultAccountService) FindAll() (*[]model.Account, *exception.Exception) {

	var err error
	var accounts *[]model.Account
	err = service.HandleTransaction(func(tx *sql.Tx) error {

		accounts, err = service.accountRepository.FindAll(tx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, exception.InternalServerErrorException(FIND_ACCOUNT_ERROR_TITLE, err)
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
