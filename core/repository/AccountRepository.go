package repository

import (
	"YellowPepper-FundsTransfers/core/model"
	"database/sql"
	"errors"
	"fmt"
)

/* TYPES DEFINITION */

type AccountRepository interface {
	Create(account *model.Account, tx *sql.Tx) error
	Update(account *model.Account, tx *sql.Tx) error
	DeleteById(id int64, tx *sql.Tx) error
	FindById(id int64, tx *sql.Tx) (*model.Account, error)
	FindAll(tx *sql.Tx) (*[]model.Account, error)
	FindByNumber(number int64, tx *sql.Tx) (*model.Account, error)
}

type DefaultAccountRepository struct {
	statementCreate       string
	statementUpdate       string
	statementDelete       string
	statementFindById     string
	statementFind         string
	statementFindByNumber string
}

/* TYPES CONSTRUCTOR */

func NewDefaultAccountRepository() *DefaultAccountRepository {
	return &DefaultAccountRepository{
		statementCreate:       "insert into account (number, balance, owner, status) values (?, ?, ?, ?)",
		statementUpdate:       "update account set number = ?, balance = ?, owner = ?, status = ? where id = ?",
		statementDelete:       "delete from account where id = ?",
		statementFindById:     "select id, number, balance, owner, status from account where id = ?",
		statementFind:         "select id, number, balance, owner, status from account",
		statementFindByNumber: "select id, number, balance, owner, status from account where number = ?",
	}
}

/* DefaultAccountRepository METHODS */

func (repository *DefaultAccountRepository) Create(account *model.Account, tx *sql.Tx) error {

	statement, err := tx.Prepare(repository.statementCreate)
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(account.Number, account.Balance, account.Owner, account.Status)
	if err != nil {
		return err
	}

	account.Id, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (repository *DefaultAccountRepository) Update(account *model.Account, tx *sql.Tx) error {

	statement, err := tx.Prepare(repository.statementUpdate)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(account.Number, account.Balance, account.Owner, account.Status, account.Id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *DefaultAccountRepository) DeleteById(id int64, tx *sql.Tx) error {

	statement, err := tx.Prepare(repository.statementDelete)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *DefaultAccountRepository) FindById(id int64, tx *sql.Tx) (*model.Account, error) {

	statement, err := tx.Prepare(repository.statementFindById)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	row := statement.QueryRow(id)

	var account model.Account
	if err := row.Scan(&account.Id, &account.Number, &account.Balance, &account.Owner, &account.Status); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New(fmt.Sprintf("account with id %d not found", id))
		}
		return nil, err
	}

	return &account, nil
}

func (repository *DefaultAccountRepository) FindAll(tx *sql.Tx) (*[]model.Account, error) {

	statement, err := tx.Prepare(repository.statementFind)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query()
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	accounts := make([]model.Account, 0)
	for rows.Next() {

		var account model.Account
		if err := rows.Scan(&account.Id, &account.Number, &account.Balance, &account.Owner, &account.Status); err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return &accounts, nil
}

func (repository *DefaultAccountRepository) FindByNumber(number int64, tx *sql.Tx) (*model.Account, error) {

	statement, err := tx.Prepare(repository.statementFindByNumber)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	row := statement.QueryRow(number)

	var account model.Account
	if err := row.Scan(&account.Id, &account.Number, &account.Balance, &account.Owner, &account.Status); err != nil {
		if err.Error() != "sql: no rows in result set" {
			return nil, err
		}
	}

	return &account, nil
}
